package fileScanner

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/Scherka/fs/tree/server/fs/server/subtypes"
)

// GetFillSize - получение размера обрабатываемой директории
func GetFillSize() {
	for _, element := range subtypes.ResponseBody.Data {
		subtypes.ResponseBody.Full_size += element.Size
	}
}

// convertSize - конвертация размеров из байт
func convertSize(size int64) string {
	prefixes := []string{"byte", "Kbyte", "Mbyte", "Gbyte", "Tbyte"}
	i := 0
	sizeFloat := float64(size)
	for (sizeFloat > subtypes.MemoryBase) && (i < 4) {
		sizeFloat = sizeFloat / subtypes.MemoryBase
		i++
	}
	return fmt.Sprintf("%.2f %s", sizeFloat, prefixes[i])
}

// sortListOfEntities - сортировка списка сущностей
func sortListOfEntities(listOfEntities []subtypes.EntityStruct, flag string) []subtypes.EntityStruct {
	if flag == subtypes.Desc {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].Size > listOfEntities[j].Size })
	} else if flag == subtypes.Asc {
		sort.Slice(listOfEntities, func(i, j int) bool { return listOfEntities[i].Size < listOfEntities[j].Size })
	}
	return listOfEntities
}

// FormatDir - добавление к root "/", если его нет
func FormatDir(dirWithoutSuffix string) string {
	if !(strings.HasSuffix(dirWithoutSuffix, "/")) {
		//приведение dir к нужному формату
		return fmt.Sprintf("%s/", dirWithoutSuffix)
	} else {
		return dirWithoutSuffix
	}
}

// getEntityParameters - получить имя, размер и тип папки/файла
func getEntityParameters(path string) (subtypes.EntityStruct, error) {
	var entity subtypes.EntityStruct
	file, err := os.Lstat(path)
	if err != nil {
		return entity, fmt.Errorf("ошибка при получении параметров %s: %v", path, err)
	}
	//если директория,то рекурсивно обходим всё её содержимое для получения размера
	if file.IsDir() {
		entity.EntityType = "Дир"
		tempSize, err := getSizeOfDir(path)
		if err != nil {
			fmt.Printf("ошибка при чтении параметров директории %s :%v\r\n", file.Name(), err)

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
func getSizeOfDir(path string) (int64, error) {
	var sizeOfDir int64
	entities, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("ошибка при чтении каталога %s: %v", path, err)
	}
	for _, entity := range entities {
		//дополняем текущий путь новым файлом/папкой
		fullPath := fmt.Sprintf("%s%s", FormatDir(path), entity.Name())
		fileStat, err := os.Lstat(fullPath)
		if err != nil {
			fmt.Printf("ошибка при получении параметров %s: %v", path, err)
		} else if fileStat.IsDir() {
			//если папка, то получаем её размер
			tempSize, err := getSizeOfDir(fullPath)
			if err != nil {
				fmt.Printf("ошибка при чтении парметров %s :%v\r\n", entity.Name(), err)
				sizeOfDir = subtypes.Multiplier * subtypes.MemoryBase
			}
			sizeOfDir += tempSize
		} else {
			sizeOfDir += fileStat.Size()

		}
	}
	return sizeOfDir, nil
}

// getListOfEntitiesParameters - получение списка папок/файлов и их свойств в корневом катлоге
func GetListOfEntitiesParameters(root string, sort string) ([]subtypes.EntityStruct, error) {
	root = FormatDir(root)
	entities, err := os.ReadDir(root)
	listOfEntitiesParameters := make([]subtypes.EntityStruct, len(entities))
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении каталога %s: %v", root, err)
	}
	//создаём группу ожидания
	wg := sync.WaitGroup{}
	for i, entity := range entities {
		wg.Add(1)
		go func(root string, listOfEntitiesParameters []subtypes.EntityStruct, entity fs.DirEntry, i int) {
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
	return sortListOfEntities(listOfEntitiesParameters, sort), nil
}
