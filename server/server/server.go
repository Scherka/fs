package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Scherka/fs/tree/server/fs/server/config"
	"github.com/Scherka/fs/tree/server/fs/server/fileScanner"
	"github.com/Scherka/fs/tree/server/fs/server/subtypes"
)

// ServerStart - запуск сервера
func ServerStart() {
	err := config.EnvParameters()
	if err != nil {
		panic(fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
	}
	server := &http.Server{
		Addr: subtypes.ConfigParam.Port,
	}

	http.HandleFunc("/fs", funcHandler)
	http.Handle("/", http.FileServer(http.Dir("./static/bundle")))
	fmt.Printf("Сервер запускается на порте: %s ", subtypes.ConfigParam.Port)
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Ошибка сервера: %v", err))
		}
	}()
	//канал сигналов
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP, syscall.SIGQUIT)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), subtypes.Multiplier*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		panic(fmt.Sprintf("Ошибка остановки сервера: %v", err))
	}
	fmt.Println("Сервер остановлен.")
}

// validateFlags - проврека флагов
func validateFlags(root, sort string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("по данному пути ничего нет")
	}
	if sort != subtypes.Asc && sort != subtypes.Desc {
		return fmt.Errorf("некорректный парметр сортировки")
	}
	return nil
}

// funcHandler - обработчик функций
func funcHandler(res http.ResponseWriter, req *http.Request) {
	subtypes.ClearResponse()
	//получение списка объектов
	root := req.FormValue("root")
	if root == "" {
		root = subtypes.ConfigParam.Root
	}
	subtypes.ResponseBody.Root = root
	sort := req.FormValue("sort")
	err := validateFlags(root, sort)
	if err != nil {
		writeErrorRespons(fmt.Sprintf("%v", err))
	} else {
		listOfEntities, err := fileScanner.GetListOfEntitiesParameters(root, sort)
		if err != nil {
			writeErrorRespons(fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		} else {
			subtypes.ResponseBody.Data = listOfEntities
			fileScanner.GetFillSize()
		}
	}
	//создание json-файла
	fileJSON, err := json.MarshalIndent(subtypes.ResponseBody, "", " ")
	if err != nil {
		writeErrorRespons(fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		//subtypes.ResponseBody.ErrorCode = 1
		//subtypes.ResponseBody.ErrorMessage = fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err)
	}
	res.Header().Set("Content-Type", "application/json")
	//вывод json на страницу
	res.Write(fileJSON)
	// Send POST request to the PHP server
	resp, err := http.Post("http://localhost/index.php", "application/json", bytes.NewBuffer(fileJSON))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(body)
}

func writeErrorRespons(error string) {
	subtypes.ResponseBody.ErrorCode = 1
	subtypes.ResponseBody.ErrorMessage = error
}
