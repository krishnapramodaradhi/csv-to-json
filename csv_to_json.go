package csv_to_json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	userInput "github.com/krishnapramodaradhi/csv-to-json/internal/user_input"
	"github.com/krishnapramodaradhi/csv-to-json/internal/util"
)

type outputDataType map[string]any

type csvToJsonJob struct {
	inputPath      string
	outputPath     string
	fileName       string
	outputFilePath string
	createdAt      time.Time
}

func (job *csvToJsonJob) GetUserInputData() error {
	reader := bufio.NewReader(os.Stdin)
	inputPath, outputPath, fileName, err := userInput.FetchData(reader)
	if err != nil {
		return err
	}
	err = userInput.ValidateData(inputPath, outputPath, fileName)
	if err != nil {
		return err
	}
	job.inputPath = inputPath
	job.outputPath = outputPath
	job.fileName = fileName
	job.outputFilePath = job.outputPath + fileName + ".json"
	return nil
}

func (job *csvToJsonJob) Process() ([]byte, error) {
	file, err := os.Open(job.inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	objects := prepareDataToConvert(scanner)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(objects)
	if err != nil {
		return nil, err
	}
	return byteData, nil
}

func (job *csvToJsonJob) WriteDataToFile(data []byte) error {
	if util.IsFileExists(job.outputFilePath) {
		return fmt.Errorf("filepath %v already exists", job.outputFilePath)
	}
	err := os.WriteFile(job.outputFilePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func New() *csvToJsonJob {
	return &csvToJsonJob{
		createdAt: time.Now(),
	}
}

func prepareDataToConvert(scanner *bufio.Scanner) []outputDataType {
	header := createHeaderList(scanner)
	objects := make([]outputDataType, 0, 20)
	for scanner.Scan() {
		row := scanner.Text()
		cells := strings.Split(row, "\t")
		newObject := make(map[string]any)
		for cellIndex, cell := range cells {
			newObject[header[cellIndex]] = cell
		}
		objects = append(objects, newObject)
	}
	return objects
}

func createHeaderList(scanner *bufio.Scanner) []string {
	scanner.Scan()
	headerString := scanner.Text()
	return strings.Split(headerString, "\t")
}
