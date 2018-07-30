package main

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
	"fmt"
)

func readCsvWithSlidingWindow(filePath string, window SlidingWindow, outputPath string){
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var time = 100
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		var value = strings.TrimSuffix(line[0], "\r\r")

		var i, conversionError = strconv.Atoi(value)

		if conversionError == nil {

			window.addDelay(time, i)
			time += 100
			var median = window.getMedian()
			var output = fmt.Sprintf("%d", median)
			var lines = []string{output}
			writeToCsv(outputPath, lines)

		} else {
			log.Fatal(conversionError)
		}
	}

	return
}

func writeToCsv(fileName string, value []string){
		f, err := getFileToWriteTo(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		w := csv.NewWriter(f)
		w.Write(value)
		w.Flush()
}

func getFileToWriteTo(fileName string)(file *os.File, err error){
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		return file, err
	} else {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		return file, err
	}
}



