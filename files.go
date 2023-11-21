package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetFileList(directory string) []string {
	var fileList []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Extract the file name from the path
			fileName := filepath.Base(path)
			fileList = append(fileList, fileName)

			fmt.Println("File in dir: ", fileName)
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error listing and appending new files:", err)
	}

	return fileList
}

func UpdateCSVStruct(fileName string, csvData []CSVRow) []CSVRow {

	// Iterate through the slice
	inCSV := false
	for _, row := range csvData {
		if fileName == row.Column1 {
			fmt.Println("File already in list: ", fileName)
			inCSV = true
			break
		}
	}

	if !inCSV {
		fmt.Println("Add file to list: ", fileName)
		// Create a new instance of the CSVRow struct and populate it with values
		var row = CSVRow{
			Column1: fileName,
			Column2: "0",
		}
		// Append the struct to the slice
		csvData = append(csvData, row)
	}

	return csvData
}
