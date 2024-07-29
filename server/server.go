package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func ServerStart() {
	err := envParameters()
	if err != nil {
		fmt.Printf("ошибка при обработке запроса:%v\r\n", err)
	}
	http.HandleFunc("/fs", func(res http.ResponseWriter, req *http.Request) { funcHandler(res, req) })
	fmt.Println("Сервер запускается")
	err = http.ListenAndServe(os.Getenv("HTTP_PORT"), nil)
	if err != nil {
		fmt.Printf("ошибка при запуске сервера: %v", err)
	}
}
func funcHandler(res http.ResponseWriter, req *http.Request) {
	//получение списка объектов
	root := req.FormValue("root")
	sort := req.FormValue("sort")
	err := fs.checkFlags(root, sort)
	switch err {
	case nil:
		listOfEntities, err := fs.getListOfEntitiesParameters(fs.formatDir(root), res)
		if err != nil {
			io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		}
		//создание json-файла
		fileJSON, err := json.MarshalIndent(listOfEntities, "", " ")
		if err != nil {
			io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		}
		res.Header().Set("Content-Type", "application/json")
		//вывод json на страницу
		res.Write(fileJSON)
	default:
		io.WriteString(res, fmt.Sprintf("%v", err))
	}

}

// envParameters - получение переменной окружения из .env
func envParameters() error {
	file, err := os.Open(".env")
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла с переменными окружения: %v", err)
	}
	defer file.Close()
	var envVar envParam
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envVar.key = splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[0]
		envVar.value = splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[1]
		os.Setenv(envVar.key, envVar.value)
	}
	return nil
}

// splitEnvParam - разбиение строки из .env
func splitEnvParam(param string) []string {
	array := regexp.MustCompile("=").Split(param, -1)
	return array
}
