package main

import (
	"flag"
	"fmt"
	"os"

	"sort"
	//"strconv"
	//"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	var root string
	root, sort, err := flagParsing()
	if err != nil {
		fmt.Printf("%v \r\n", err)
		return
	}
	listOfEntities, err := getListOfEntitiesParameters(root)
	if err != nil {
		fmt.Printf("%v \r\n", err)
		return
	}
	//fmt.Println(listOfEntities)
	output(sortListOfEntities(listOfEntities, sort))
}

// entity - содержит имя, тип и размер папки/файла
type entityStruct struct {
	name       string
	entityType string
	size       int64
}

// getMaxLenOfName - определить максимальную длину имени
func getMaxLenOfName(listOfEntities []entityStruct) int {
	max := 0
	for _, e := range listOfEntities {
		if utf8.RuneCountInString(e.name) > max {
			max = utf8.RuneCountInString(e.name)
		}
	}
	//fmt.Println(max)
	return max
}

// output - вывод в нужном формате
func output(listOfEntities []entityStruct) {
	maxLen := getMaxLenOfName(listOfEntities)
	fmt.Printf("Тип%sИмя%sРазмер\r\n", strings.Repeat(" ", 2), strings.Repeat(" ", maxLen-2))
	for _, e := range listOfEntities {
		//fmt.Println(e.entityType, utf8.RuneCountInString(e.entityType))
		fmt.Printf("%s%s%s%s%s\r\n", e.entityType,
			strings.Repeat(" ", 5-utf8.RuneCountInString(e.entityType)), e.name,
			strings.Repeat(" ", maxLen-utf8.RuneCountInString(e.name)+1), convertSize(e.size))
	}
}

// convertSize - конвертация размеров из байт
func convertSize(size int64) string {
	prefixes := []string{"byte", "kbyte", "mbyte", "gbyte"}
	i := 0
	for (size > 1000) && (i < 3) {
		size = size / 1000
		i++
	}
	return fmt.Sprintf("%d %s", size, prefixes[i])
}

// sortListOfEntities - сортировка списка сущностей
func sortListOfEntities(listOfEntities []entityStruct, flag string) []entityStruct {
	if flag == "desc" {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].size > listOfEntities[j].size })
	} else if flag == "asc" {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].size < listOfEntities[j].size })
	}
	return listOfEntities
}

// flagParsing - обработка флагов
func flagParsing() (string, string, error) {

	//флаг каталога
	root := flag.String("root", "", "используйте флаг -root для введения сканируемого каталога.")
	//флаг сортировки
	sort := flag.String("sort", "", "используйте флаг -sort для введения порядка сортировки: asc - по возрастанию, desc= по убыванию.")
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
	if *sort != "asc" && *sort != "desc" {
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
	file, err := os.Stat(path)
	if err != nil {
		return entity, fmt.Errorf("ошибка при получении парсметров %s: %v", path, err)
	}
	//если директория,то рекурсивно обходим всё её содержимое для получения размера
	if file.IsDir() {
		entity.entityType = "Дир"
		tempSize, err := getSizeOfDir(path)
		if err != nil {
			fmt.Printf("ошибка при чтении парметров %s :%v\r\n", file.Name(), err)

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
	for _, e := range entities {
		//дополняем текущий путь новым файлом/папкой
		fullPath := fmt.Sprintf("%s%s", formatDir(path), e.Name())
		fileStat, err := os.Lstat(fullPath)
		if err != nil {
			return 0, fmt.Errorf("ошибка при получении парсметров %s: %v", path, err)
		}
		if fileStat.IsDir() {
			//если папка, то получаем её размер
			tempSize, err := getSizeOfDir(fullPath)
			if err != nil {
				fmt.Printf("ошибка при чтении парметров %s :%v\r\n", e.Name(), err)

			}
			sizeOfDir += tempSize
		} else {
			sizeOfDir += fileStat.Size()

		}
	}
	return sizeOfDir, nil
}

// getListOfEntities - получение списка папока и файлов в корневом катлоге
func getListOfEntitiesParameters(root string) ([]entityStruct, error) {
	listOfEntitiesParameters := make([]entityStruct, 0)
	entities, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении каталога %s: %v", root, err)
	}
	for _, e := range entities {
		entityParameters, err := getEntityParameters(fmt.Sprintf("%s%s", root, e.Name()))
		if err != nil {
			fmt.Printf("ошибка при чтении парметров %s :%v\r\n", e.Name(), err)
		} else {
			listOfEntitiesParameters = append(listOfEntitiesParameters, entityParameters)
		}
		//fmt.Println(getEntityParameters(fmt.Sprintf("%s%s", root, e.Name())))
		//fmt.Println(e.Name(), e.IsDir())
	}
	return listOfEntitiesParameters, nil
}
