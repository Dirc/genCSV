package main

import (
	"github.com/alecthomas/kingpin/v2"
)

var (
	dir     = kingpin.Flag("dir", "Dir to list files").Default(".").String()
	csvFile = kingpin.Flag("csv", "CSV file to create or update").Default("files.csv").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	UpdateCSV(*dir, *csvFile)
}
