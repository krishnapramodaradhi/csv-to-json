package main

import (
	"fmt"

	csvToJson "github.com/krishnapramodaradhi/csv-to-json"
)

func main() {
	c2j := csvToJson.New()
	err := c2j.GetUserInputData()
	if err != nil {
		fmt.Println("There was an error in getting the user input", err)
		return
	}
	byteData, err := c2j.Process()
	if err != nil {
		fmt.Println("An error occured while processing the data", err)
		return
	}

	err = c2j.WriteDataToFile(byteData)
	if err != nil {
		fmt.Println("An error occured while writing the file to fs", err)
		return
	}
	fmt.Println("File write complete")
}
