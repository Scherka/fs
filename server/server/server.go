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

	http.HandleFunc("/fs", fsHandler)
	http.HandleFunc("/statistic", statHandler)
	http.Handle("/", http.FileServer(http.Dir("./static/bundle")))
	fmt.Printf("Сервер запускается на порте: %s\n", subtypes.ConfigParam.Port)
	//Горутина для graceful shutdown
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

// fsHandler - обработчик функций
func fsHandler(res http.ResponseWriter, req *http.Request) {
	start := time.Now()
	var responseBody subtypes.Response
	//получение списка объектов
	root := req.FormValue("root")
	if root == "" {
		root = subtypes.ConfigParam.Root
	}
	responseBody.Root = root
	sort := req.FormValue("sort")
	err := validateFlags(root, sort)
	if err != nil {
		writeErrorRespons(&responseBody, fmt.Sprintf("%v", err))
	} else {
		listOfEntities, err := fileScanner.GetListOfEntitiesParameters(root, sort)
		if err != nil {
			writeErrorRespons(&responseBody, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
		} else {
			responseBody.Data = listOfEntities
			//получение полного размера директории
			responseBody.FullSize = fileScanner.GetFullSize(responseBody)
		}
	}
	//вычисление времени работы сканера и запись в тело ответа
	finish := time.Now()
	responseBody.DateOfRequest = finish.Format("2006-01-02")
	responseBody.TimeOfRequest = finish.Format("15:04:05")
	responseBody.LoadingTime = float64(time.Since(start).Truncate(100 * time.Microsecond))
	fileJSON, err := json.MarshalIndent(responseBody, "", " ")
	if err != nil {
		writeErrorRespons(&responseBody, fmt.Sprintf("ошибка при обработке запроса:%v\r\n", err))
	}
	res.Header().Set("Content-Type", "application/json")
	//вывод json на страницу
	res.Write(fileJSON)
	// POST - запрос
	resp, err := http.Post(subtypes.ConfigParam.DB_INSERTER_PATH, "application/json", bytes.NewBuffer(fileJSON))
	if err != nil {
		fmt.Printf("Ошибка при выполнении POST-запроса:  %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при выполнении POST-запроса %v:  %v", resp.Status, err)
		return
	}
	fmt.Printf("%s %s", resp.Status, string(body))
}
func statHandler(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(subtypes.ConfigParam.STAT_DISPLAY_PATH)
	if err != nil {
		fmt.Printf("Ошибка при выполнении GET-запроса:  %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при выполнении GET-запроса:  %v", err)
		return
	}
	res.Header().Set("Content-Type", "text/html")
	res.Write(body)
}
func writeErrorRespons(resp *subtypes.Response, error string) {
	resp.ErrorCode = 1
	resp.ErrorMessage = error
	finish := time.Now()
	resp.DateOfRequest = finish.Format("02-01-2006")
	resp.TimeOfRequest = finish.Format("15:04:05")
}
