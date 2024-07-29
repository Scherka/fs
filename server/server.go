package server

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/Scherka/fs/tree/server/fs/fileScanner"
)

// ServerStart - запуск сервера
func ServerStart() {
	err := envParameters()
	if err != nil {
		fmt.Printf("ошибка при обработке запроса:%v\r\n", err)
	}
	server := &http.Server{
		Addr: os.Getenv("HTTP_PORT"),
	}

	http.HandleFunc("/fs", funcHandler)
	fmt.Println("Сервер запускается")
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Ошибка сервера: %v", err)
		}
	}()
	//канал сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGKILL, syscall.SIGQUIT)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Ошибка остановки сервера: %v", err)
	}
	fmt.Println("Сервер остановлен.")
}

// checkFlags - проврека флагов
func checkFlags(root, sort string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("Объекта не существует")
	}
	if sort != fileScanner.Asc && sort != fileScanner.Desc {
		return fmt.Errorf("Некорректный парметр сортировки")
	}
	return nil
}
func funcHandler(res http.ResponseWriter, req *http.Request) {
	//получение списка объектов
	root := req.FormValue("root")
	sort := req.FormValue("sort")
	err := checkFlags(root, sort)
	switch err {
	case nil:
		listOfEntities, err := fileScanner.GetListOfEntitiesParameters(root, res, sort)
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
	var envVar fileScanner.EnvParam
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envVar.Key = splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[0]
		envVar.Value = splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[1]
		os.Setenv(envVar.Key, envVar.Value)
	}
	return nil
}

// splitEnvParam - разбиение строки из .env
func splitEnvParam(param string) []string {
	array := regexp.MustCompile("=").Split(param, -1)
	return array
}
