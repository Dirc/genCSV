package main

import (
	"reflect"
	"testing"
)

func TestSetTotal(t *testing.T) {
	t.Run("UpdatesToTal", func(t *testing.T) {
		input := []CSVRow{
			{Column1: "file1", Column2: "10"},
			{Column1: "file2", Column2: "3"},
			{Column1: "total", Column2: "8"},
			{Column1: "file3", Column2: "1"},
		}

		expected := []CSVRow{
			{Column1: "file1", Column2: "10"},
			{Column1: "file2", Column2: "3"},
			{Column1: "total", Column2: "14"}, // updated value
			{Column1: "file3", Column2: "1"},
		}

		result := SetTotal(input)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %+v, Got: %+v", expected, result)
		}
	})
}
