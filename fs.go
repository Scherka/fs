package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// entitySruct - содержит имя, тип и размер папки/файла
type entityStruct struct {
	name       string //Имя объекта
	entityType string //Тип объекта
	size       int64  //Размер объекта в байтах
}

const asc = "asc"       //флаг сортировки по возрастанию
const desc = "desc"     //флаг сортировки по убыванию
const memoryBase = 1000 //основание конвертации памяти

func main() {
	//время начала программы
	start := time.Now()
	var root string
	root, sort, err := flagParsing()
	if err != nil {
		panic(fmt.Sprintf("%v \r\n", err))
	}
	listOfEntities, err := getListOfEntitiesParameters(root)
	if err != nil {
		panic(fmt.Sprintf("%v \r\n", err))
	}
	output(sortListOfEntities(listOfEntities, sort))
	//время выполнения программы
	finish := time.Since(start).Truncate(10 * time.Millisecond).String()
	fmt.Println("Время выполнения программы:", finish)
}

// getMaxLenOfName - определить максимальную длину имени
func getMaxLenOfName(listOfEntities []entityStruct) int {
	max := 0
	for _, entity := range listOfEntities {
		if utf8.RuneCountInString(entity.name) > max {
			max = utf8.RuneCountInString(entity.name)
		}
	}
	return max
}

// output - вывод в нужном формате
func output(listOfEntities []entityStruct) {
	maxLen := getMaxLenOfName(listOfEntities)
	fmt.Printf("Тип%sИмя%sРазмер\r\n", strings.Repeat(" ", 2), strings.Repeat(" ", maxLen-2))
	for _, entity := range listOfEntities {
		fmt.Printf("%s%s%s%s%s\r\n", entity.entityType,
			strings.Repeat(" ", 5-utf8.RuneCountInString(entity.entityType)), entity.name,
			strings.Repeat(" ", maxLen-utf8.RuneCountInString(entity.name)+1), convertSize(entity.size))
	}
}

// convertSize - конвертация размеров из байт
func convertSize(size int64) string {
	prefixes := []string{"byte", "kbyte", "mbyte", "gbyte", "tbyte"}
	i := 0
	sizeFloat := float64(size)
	for (sizeFloat > memoryBase) && (i < 4) {
		sizeFloat = sizeFloat / memoryBase
		i++
	}
	return fmt.Sprintf("%.2f %s", sizeFloat, prefixes[i])
}

// sortListOfEntities - сортировка списка сущностей
func sortListOfEntities(listOfEntities []entityStruct, flag string) []entityStruct {
	if flag == desc {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].size > listOfEntities[j].size })
	} else if flag == asc {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].size < listOfEntities[j].size })
	}
	return listOfEntities
}

// flagParsing - обработка флагов
func flagParsing() (string, string, error) {
	const asc = "asc"
	const desc = "desc"
	//флаг каталога
	root := flag.String("root", "", "используйте флаг -root для введения сканируемого каталога.")
	//флаг сортировки
	sort := flag.String("sort", "", "используйте флаг -sort для введения порядка сортировки: asc - по возрастанию, desc - по убыванию.")
	flag.Parse()
	//проверка наличия флагов
	if len(*root) == 0 {
		flag.PrintDefaults()
		return "", "", fmt.Errorf("отстутствуют необходимые флаги: -root")
	}
	if len(*sort) == 0 {
		flag.PrintDefaults()
		return "", "", fmt.Errorf("отстутствуют необходимые флаги: -sort")
	}
	if *sort != asc && *sort != desc {
		flag.PrintDefaults()
		return "", "", fmt.Errorf("флаг -sort задан в неверном формате")
	}
	return formatDir(*root), *sort, nil
}

// formatDir - добавление к root "/", если его нет
func formatDir(dirWithoutSuffix string) string {
	if !(strings.HasSuffix(dirWithoutSuffix, "/")) {
		//приведение dir к нужному формату
		return fmt.Sprintf("%s/", dirWithoutSuffix)
	} else {
		return dirWithoutSuffix
	}
}

// getEntityParameters - получить имя, размер и тип папки/файла
func getEntityParameters(path string) (entityStruct, error) {
	var entity entityStruct
	file, err := os.Lstat(path)
	if err != nil {
		return entity, fmt.Errorf("ошибка при получении параметров %s: %v", path, err)
	}
	//если директория,то рекурсивно обходим всё её содержимое для получения размера
	if file.IsDir() {
		entity.entityType = "Дир"
		tempSize, err := getSizeOfDir(path)
		if err != nil {
			fmt.Printf("ошибка при чтении параметров директории %s :%v\r\n", file.Name(), err)

		} else {
			entity.size += tempSize
		}
	} else {
		entity.entityType = "Файл"
		entity.size += file.Size()
	}
	entity.name = file.Name()
	return entity, nil
}

// getSizeOfDir - получение размера папки
func getSizeOfDir(path string) (int64, error) {
	var sizeOfDir int64
	entities, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении каталога %s: %v", path, err)
	}
	for _, entity := range entities {
		//дополняем текущий путь новым файлом/папкой
		fullPath := fmt.Sprintf("%s%s", formatDir(path), entity.Name())
		fileStat, err := os.Lstat(fullPath)
		if err != nil {
			fmt.Printf("ошибка при получении параметров %s: %v", path, err)
		} else if fileStat.IsDir() {
			//если папка, то получаем её размер
			tempSize, err := getSizeOfDir(fullPath)
			if err != nil {
				fmt.Printf("ошибка при чтении парметров %s :%v\r\n", entity.Name(), err)
				sizeOfDir = 4 * memoryBase
			}
			sizeOfDir += tempSize
		} else {
			sizeOfDir += fileStat.Size()

		}
	}
	return sizeOfDir, nil
}

// getListOfEntitiesParameters - получение списка папок/файлов и их свойств в корневом катлоге
func getListOfEntitiesParameters(root string) ([]entityStruct, error) {
	entities, err := os.ReadDir(root)
	listOfEntitiesParameters := make([]entityStruct, len(entities))
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении каталога %s: %v", root, err)
	}
	//создаём группу ожидания
	wg := sync.WaitGroup{}
	for i, entity := range entities {
		wg.Add(1)
		go func(root string, listOfEntitiesParameters []entityStruct, entity fs.DirEntry, i int) {
			defer wg.Done()
			//получаем параметры объекта
			entityParameters, err := getEntityParameters(fmt.Sprintf("%s%s", root, entity.Name()))
			if err != nil {
				fmt.Printf("ошибка при чтении параметров %s :%v\r\n", entity.Name(), err)
			} else {
				listOfEntitiesParameters[i] = entityParameters
			}
		}(root, listOfEntitiesParameters, entity, i)
	}
	wg.Wait()
	return listOfEntitiesParameters, nil
}
