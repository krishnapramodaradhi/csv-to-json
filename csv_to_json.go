package csvToJson

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	userInput "github.com/krishnapramodaradhi/csv-to-json/internal/userInput"
	"github.com/krishnapramodaradhi/csv-to-json/internal/util"
)

type outputJson map[string]any

type csvToJson struct {
	inputPath      string
	outputPath     string
	fileName       string
	outputFilePath string
	fileExtension  string
	createdAt      time.Time
}

func (c *csvToJson) GetUserInputData() error {
	reader := bufio.NewReader(os.Stdin)
	inputPath, outputPath, fileName, err := userInput.FetchData(reader)
	if err != nil {
		return err
	}
	err = userInput.ValidateData(inputPath, outputPath, fileName)
	if err != nil {
		return err
	}
	c.inputPath = inputPath
	c.outputPath = outputPath
	c.fileName = fileName
	c.fileExtension = ".json"
	c.outputFilePath = c.outputPath + c.fileName + c.fileExtension
	return nil
}

func (c *csvToJson) Process() ([]byte, error) {
	file, err := os.Open(c.inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	objects, err := prepareDataToConvert(scanner)
	if err != nil {
		return nil, err
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(objects)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}

func (c *csvToJson) WriteDataToFile(data []byte) error {
	if util.IsFileExists(c.outputFilePath) {
		return fmt.Errorf("filepath %v already exists", c.outputFilePath)
	}
	err := os.WriteFile(c.outputFilePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func New() *csvToJson {
	return &csvToJson{
		createdAt: time.Now(),
	}
}

func prepareDataToConvert(scanner *bufio.Scanner) ([]outputJson, error) {
	header, delimiter := createHeaderList(scanner)
	if header == nil {
		return nil, fmt.Errorf("no supported delimiter (Comma, Pipe or tab space) found in the file")
	}
	objects := make([]outputJson, 0, 20)
	for scanner.Scan() {
		row := scanner.Text()
		cells := strings.Split(row, delimiter)
		newObject := make(map[string]any)
		for cellIndex, cell := range cells {
			newObject[header[cellIndex]] = cell
		}
		objects = append(objects, newObject)
	}
	return objects, nil
}

var supportedDelimiters = [3]string{",", "|", "\t"}

func createHeaderList(scanner *bufio.Scanner) ([]string, string) {
	scanner.Scan()
	headerString := scanner.Text()
	for _, delimiter := range supportedDelimiters {
		if strings.Contains(headerString, delimiter) {
			return strings.Split(headerString, delimiter), delimiter
		}
	}
	return nil, ""
}
