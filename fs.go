package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// entitySruct - содержит имя, тип и размер папки/файла
type entityStruct struct {
	Name          string `json:"Имя"`    //Имя объекта
	EntityType    string `json:"Тип"`    //Тип объекта
	Size          int64  `json:"-"`      //Размер объекта в байтах
	SizeFormatted string `json:"Размер"` //Форматированный размер объекта
}

const asc = "asc"       //флаг сортировки по возрастанию
const desc = "desc"     //флаг сортировки по убыванию
const memoryBase = 1000 //основание конвертации памяти

func main() {
	//https://localhost/fs?root=/home/sergey&sort=asc
	envParameters()
	http.HandleFunc("/fs", func(res http.ResponseWriter, req *http.Request) { handler(res, req) })
	http.ListenAndServe(os.Getenv("HTTP_PORT"), nil)

}

// envParameters - получение переменной окружения из port.env
func envParameters() {
	file, _ := os.Open("port.env")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		os.Setenv(splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[0],
			splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[1])
	}
}

// splitEnvParam - разбиение строки из .env
func splitEnvParam(param string) []string {
	array := regexp.MustCompile("=").Split(param, -1)
	return array
}
func handler(res http.ResponseWriter, req *http.Request) {
	//время начала программы
	start := time.Now()
	//получение списка объектов
	root := req.FormValue("root")
	sort := req.FormValue("sort")
	err := checkFlags(root, sort)
	switch err {
	case nil:
		listOfEntities, err := getListOfEntitiesParameters(formatDir(root), res)
		if err != nil {
			io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		}
		//создание json-файла
		fileJSON, err := encodeJSON(sortListOfEntities(listOfEntities, req.FormValue("sort")))
		if err != nil {
			io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		}
		//вывод json на страницу
		io.WriteString(res, fmt.Sprintf("%s\n", string(fileJSON)))
		//время выполнения программы
		finish := time.Since(start).Truncate(10 * time.Millisecond).String()
		io.WriteString(res, fmt.Sprintf("Время выполнения завроса: %s", finish))
	default:
		io.WriteString(res, fmt.Sprintf("%v", err))
	}
}

// checkFlags - проврека флагов
func checkFlags(root, sort string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Объекта не существует")
	}
	if sort != asc && sort != desc {
		return fmt.Errorf("Некорректный парметр сортировки")
	}
	return nil
}

// encodeJSON - создать json-файл
func encodeJSON(listOfEntities []entityStruct) ([]byte, error) {
	fileJSON, err := json.MarshalIndent(listOfEntities, "", " ")
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании JSON-файла:%v", err)
	}
	return fileJSON, nil

}

// convertSize - конвертация размеров из байт
func convertSize(size int64) string {
	prefixes := []string{"byte", "Kbyte", "Mbyte", "Gbyte", "Tbyte"}
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
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].Size > listOfEntities[j].Size })
	} else if flag == asc {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].Size < listOfEntities[j].Size })
	}
	return listOfEntities
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
func getEntityParameters(path string, res http.ResponseWriter) (entityStruct, error) {
	var entity entityStruct
	file, err := os.Lstat(path)
	if err != nil {
		return entity, fmt.Errorf("ошибка при получении параметров %s: %v", path, err)
	}
	//если директория,то рекурсивно обходим всё её содержимое для получения размера
	if file.IsDir() {
		entity.EntityType = "Дир"
		tempSize, err := getSizeOfDir(path, res)
		if err != nil {
			io.WriteString(res, fmt.Sprintf("ошибка при чтении параметров директории %s :%v\r\n", file.Name(), err))

		} else {
			entity.Size += tempSize
		}
	} else {
		entity.EntityType = "Файл"
		entity.Size += file.Size()
	}
	entity.Name = file.Name()
	entity.SizeFormatted = convertSize(entity.Size)
	return entity, nil
}

// getSizeOfDir - получение размера папки
func getSizeOfDir(path string, res http.ResponseWriter) (int64, error) {
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
			io.WriteString(res, fmt.Sprintf("ошибка при получении параметров %s: %v", path, err))
		} else if fileStat.IsDir() {
			//если папка, то получаем её размер
			tempSize, err := getSizeOfDir(fullPath, res)
			if err != nil {
				io.WriteString(res, fmt.Sprintf("ошибка при чтении парметров %s :%v\r\n", entity.Name(), err))
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
func getListOfEntitiesParameters(root string, res http.ResponseWriter) ([]entityStruct, error) {
	entities, err := os.ReadDir(root)
	listOfEntitiesParameters := make([]entityStruct, len(entities))
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении каталога %s: %v", root, err)
	}
	//создаём группу ожидания
	wg := sync.WaitGroup{}
	for i, entity := range entities {
		wg.Add(1)
		go func(root string, listOfEntitiesParameters []entityStruct, entity fs.DirEntry, i int, res http.ResponseWriter) {
			defer wg.Done()
			//получаем параметры объекта
			entityParameters, err := getEntityParameters(fmt.Sprintf("%s%s", root, entity.Name()), res)
			if err != nil {
				io.WriteString(res, fmt.Sprintf("ошибка при чтении параметров %s :%v\r\n", entity.Name(), err))
			} else {
				listOfEntitiesParameters[i] = entityParameters
			}
		}(root, listOfEntitiesParameters, entity, i, res)
	}
	wg.Wait()
	return listOfEntitiesParameters, nil
}
