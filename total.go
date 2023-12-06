package main

import (
	"log"
	"strconv"
)

func SetTotal(csvData []CSVRow) []CSVRow {

	// Calculate total
	var total int
	for _, row := range csvData {
		if row.Column1 != "total" {
			price, err := strconv.Atoi(row.Column2)
			if err != nil {
				panic(err)
			}
			total += price
			log.Printf("Add: %s, new total: %v\n", row.Column2, total)
		}
	}
	totalString := strconv.Itoa(total)

	var totalExists = false
	// if row "total" exists, update it
	for i := range csvData {
		if csvData[i].Column1 == "total" {
			csvData[i].Column2 = totalString
			totalExists = true
		}
	}

	// if row "total" doesn't exists, create it
	if !totalExists {
		var row = CSVRow{
			Column1: "total",
			Column2: totalString,
		} // Append the struct to the slice
		csvData = append(csvData, row)
	}

	log.Printf("Total: %v\n", total)

	return csvData
}
