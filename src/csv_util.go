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
			var output = fmt.Sprintf("%d \r\r", median)
			var lines = []string{output}
			writeToCsv(outputPath, lines)

		} else {
			log.Fatal(conversionError)
		}
	}

	return
}

func writeToCsv(fileName string, value []string){
	file, err := os.Create(fileName)
	if err == nil {
		writer := csv.NewWriter(file)
		defer writer.Flush()

		defer file.Close()

		defer writer.Flush()

		err = writer.Write(value)

		if err != nil {
			file.Close()
			log.Fatal(err)
		}
	}
}



