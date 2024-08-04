package utils

import (
	"fmt"
	"os"
)

func FileAsString(name string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/inputs/%s", cwd, name)
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	fileString := string(file)
	return fileString, nil
}
