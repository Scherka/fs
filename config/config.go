package config

import (
	"bufio"
	"fmt"

	"os"

	"regexp"
	"strings"

	"github.com/Scherka/fs/tree/server/fs/fileScanner"
	"github.com/Scherka/fs/tree/server/fs/subtypes"
)

// EnvParameters - получение переменной окружения из .env
func EnvParameters() error {
	file, err := os.Open(".env")
	if err != nil {
		return fmt.Errorf("ошибка при открытии файла с переменными окружения: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[0] {
		case "HTTP_PORT":
			subtypes.ConfigParam.Port = splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[1]
		case "ROOT":
			subtypes.ConfigParam.Root = fileScanner.FormatDir(splitEnvParam(strings.ReplaceAll(scanner.Text(), " ", ""))[1])
		}

	}
	return nil
}

// splitEnvParam - разбиение строки из .env
func splitEnvParam(param string) []string {
	array := regexp.MustCompile("=").Split(param, -1)
	return array
}
