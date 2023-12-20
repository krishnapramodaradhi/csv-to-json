package main

import (
	"fmt"

	csv_to_json "github.com/krishnapramodaradhi/csv-to-json"
)

func main() {
	csvToJson := csv_to_json.New()
	err := csvToJson.GetUserInputData()
	if err != nil {
		fmt.Println("There was an error in getting the user input", err)
		return
	}
	byteData, err := csvToJson.Process()
	if err != nil {
		fmt.Println("An error occured while processing the data", err)
		return
	}

	err = csvToJson.WriteDataToFile(byteData)
	if err != nil {
		fmt.Println("An error occured while writing the file to fs", err)
		return
	}
	fmt.Println("File write complete")
}
