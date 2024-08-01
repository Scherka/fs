package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/Scherka/fs/tree/server/fs/config"
	"github.com/Scherka/fs/tree/server/fs/fileScanner"
	"github.com/Scherka/fs/tree/server/fs/subtypes"
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
	http.Handle("/bundle/", http.StripPrefix("/bundle", http.FileServer(http.Dir("./bundle"))))
	http.HandleFunc("/", StartPage)
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

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
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
	switch err {
	case nil:
		listOfEntities, err := fileScanner.GetListOfEntitiesParameters(root, sort)
		if err != nil {
			//io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
			subtypes.ResponseBody.ErrorCode = 1
			subtypes.ResponseBody.ErrorMessage = fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err)
		} else {
			subtypes.ResponseBody.Data = listOfEntities
		}

	default:
		//io.WriteString(res, fmt.Sprintf("%v", err))
		subtypes.ResponseBody.ErrorCode = 1
		subtypes.ResponseBody.ErrorMessage = fmt.Sprint(err)
	}
	//создание json-файла
	fileJSON, err := json.MarshalIndent(subtypes.ResponseBody, "", " ")
	if err != nil {
		//io.WriteString(res, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		subtypes.ResponseBody.ErrorCode = 1
		subtypes.ResponseBody.ErrorMessage = fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err)
	}
	res.Header().Set("Content-Type", "application/json")
	//вывод json на страницу
	res.Write(fileJSON)
}

func StartPage(rw http.ResponseWriter, r *http.Request) {
	//создаем html-шаблон
	tmpl, err := template.ParseFiles("./bundle/bundle.html")
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//выводим шаблон клиенту в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
