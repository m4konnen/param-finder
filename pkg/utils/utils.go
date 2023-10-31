package utils

import (
	"bufio"
	"errors"
	"os"
)

// Convert fileLines to an array
func FileToList(filename string) ([]string, error) {

	var newList []string

	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		newList = append(newList, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("error Reading file")
	}

	return newList, nil
}

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
