package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type CSVRow struct {
	Column1 string
	Column2 string
}

func CSVToStruct(filename string) []CSVRow {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
	}

	// Create a slice to store the parsed data
	var csvData []CSVRow

	// Iterate through the records
	for _, record := range records {
		// Assuming there are two columns in each record
		if len(record) == 2 {
			// Create a new instance of the CSVRow struct and populate it with values
			var row = CSVRow{
				Column1: record[0],
				Column2: record[1],
			} // Append the struct to the slice
			csvData = append(csvData, row)
		} else {
			fmt.Println("Invalid record:", record)
		}
	}

	// Now you have a slice of CSVRow structs representing your CSV data
	// You can iterate through csvData and make changes as needed
	for _, row := range csvData {
		fmt.Printf("Column1: %s, Column2: %s\n", row.Column1, row.Column2)
	}

	return csvData
}

func StructToCSV(csvData []CSVRow, filename string) {

	// Create a CSV file
	file, err := os.Create(filename) // Replace with your desired output file name
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write data
	for _, row := range csvData {
		writer.Write([]string{row.Column1, row.Column2})
	}

	fmt.Printf("CSV file created successfully: %s", filename)
}

func UpdateCSV(directory string, csvFileName string) {
	fullPath := filepath.Join(directory, csvFileName)

	csvData := CSVToStruct(fullPath)
	fileList := GetFileList(directory)
	fmt.Println("fileList: ", fileList)

	for _, fileName := range fileList {
		csvData = UpdateCSVStruct(fileName, csvData)
	}

	csvData = SetTotal(csvData)
	StructToCSV(csvData, fullPath)

}
