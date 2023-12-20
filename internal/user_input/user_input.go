package userInput

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/krishnapramodaradhi/csv-to-json/internal/util"
)

func FetchData(reader *bufio.Reader) (inputPath, outputPath, fileName string, err error) {
	inputPath, err = readUserInput("Path to the TSV file: ", reader)
	if err != nil {
		return "", "", "", err
	}
	outputPath, err = readUserInput("Output Path: ", reader)
	if err != nil {
		return "", "", "", err
	}
	fileName, err = readUserInput("File Name: ", reader)
	if err != nil {
		return "", "", "", err
	}
	return
}

func ValidateData(inputPath, outputPath, fileName string) error {
	if !util.IsFileExists(inputPath) {
		return errors.New("input path does not exist")
	}
	if !util.IsFileExists(outputPath) {
		return errors.New("output path does not exist")
	}
	if len(fileName) == 0 {
		return errors.New("invalid Filename")
	}
	return nil
}

func readUserInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value, nil
}
