package main

import (
	"reflect"
	"testing"
)

func TestUpdateCSV(t *testing.T) {
	t.Run("UpdateCSV", func(t *testing.T) {
		input := []CSVRow{
			{Column1: "file1", Column2: "10"},
			{Column1: "file2", Column2: "3"},
		}

		fileName := "file3"

		expected := []CSVRow{
			{Column1: "file1", Column2: "10"},
			{Column1: "file2", Column2: "3"},
			{Column1: "file3", Column2: "0"},
		}

		result := UpdateCSVStruct(fileName, input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, Got: %+v", expected, result)
		}
	})
}
