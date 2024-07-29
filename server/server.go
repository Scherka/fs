package server

import (
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

	"github.com/Scherka/fs/tree/server/fs/config"
	"github.com/Scherka/fs/tree/server/fs/fileScanner"
	"github.com/Scherka/fs/tree/server/fs/subtypes"
)

// ServerStart - запуск сервера
func ServerStart() {
	err := config.EnvParameters()
	if err != nil {
		fmt.Printf("ошибка при обработке запроса:%v\r\n", err)
	}
	server := &http.Server{
		Addr: subtypes.Port.Value,
	}

	http.HandleFunc("/fs", funcHandler)
	fmt.Printf("Сервер запускается на порте: %s", subtypes.Port.Value)
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

// validateFlags - проврека флагов
func validateFlags(root, sort string) error {
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("По данному пути ничего нет")
	}
	if sort != subtypes.Asc && sort != subtypes.Desc {
		return fmt.Errorf("Некорректный парметр сортировки")
	}
	return nil
}

// funcHandler - обработчик функций
func funcHandler(res http.ResponseWriter, req *http.Request) {
	//получение списка объектов
	root := req.FormValue("root")
	sort := req.FormValue("sort")
	err := validateFlags(root, sort)
	switch err {
	case nil:
		listOfEntities, err := fileScanner.GetListOfEntitiesParameters(root, sort)
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
