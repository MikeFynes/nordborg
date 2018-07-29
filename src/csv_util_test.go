package main

import (
	"testing"
	"os"
)

func TestCsvReaderDoc1(t *testing.T){
	var window = SlidingWindow{}
	window.size = 5
	window.timeDiffMax = 25000
	var outputFile = "output1.csv"
	if _, err := os.Stat(outputFile); !os.IsNotExist(err) {
		var err = os.Remove(outputFile)
		if err != nil {
			t.Errorf("could not delete file %s", outputFile)
		}
	}

	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test2.csv", window, outputFile)
}

func TestCsvReaderDoc2(t *testing.T) {
	var window= SlidingWindow{}
	window.size = 5
	window.timeDiffMax = 25000
	var outputFile = "output2.csv"
	if _, err := os.Stat(outputFile); os.IsExist(err) {
		var err = os.Remove(outputFile)
		if err != nil {
			t.Errorf("could not delete file %s", outputFile)
		}
	}
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test3.csv", window, outputFile)
}

func TestCsvReaderDoc3(t *testing.T){
	var window = SlidingWindow{}
	window.size = 5
	window.timeDiffMax = 25000
	var outputFile = "output3.csv"
	if _, err := os.Stat(outputFile); os.IsExist(err) {
		var err = os.Remove(outputFile)
		if err != nil {
			t.Errorf("could not delete file %s", outputFile)
		}
	}
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test4.csv", window, outputFile)
}
