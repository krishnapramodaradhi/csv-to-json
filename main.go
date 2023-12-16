package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Product struct {
	Id               string  `json:"id"`
	Title            string  `json:"title"`
	ShortDescription string  `json:"shortDescription"`
	Description      string  `json:"description"`
	Category         string  `json:"category"`
	Price            float64 `json:"price"`
	Quantity         int     `json:"quantity"`
	ImageUrl         string  `json:"imageUrl"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Path to CSV file: ")
	inputPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An issue occured while reading the path", err)
		return
	}
	inputPath = strings.TrimSuffix(inputPath, "\n")
	inputPath = strings.TrimSuffix(inputPath, "\r")
	fmt.Print("Output Dir: ")
	outputPath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An issue occured while reading the path", err)
		return
	}
	outputPath = strings.TrimSuffix(outputPath, "\n")
	outputPath = strings.TrimSuffix(outputPath, "\r")
	validPath := regexp.MustCompile(`^(?:[\w]\:|\\)(\\[a-zA-Z_\-\s0-9\.]+)+\.(csv)`)
	match := validPath.MatchString(inputPath)
	if !match {
		fmt.Println("The input path is invalid")
		return
	}

	match = validPath.MatchString(outputPath)
	if !match {
		fmt.Println("The input path is invalid")
		return
	}

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("An error occured while reading the file", err)
		return
	}
	defer file.Close()

	// products := make([]Product, 0, 20)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanner.Text()
	}
}
