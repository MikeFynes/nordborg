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

func readCsvWithSlidingWindow(filePath string, maxWindowSize int, maxTimeSeparationMillis int)(values []int){
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var time = 100
	var window = SlidingWindow{}
	window.size = maxWindowSize
	window.timeDiffMax = maxTimeSeparationMillis
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
			fmt.Println(median)
			values = append(values, i)
		} else {
			log.Fatal(conversionError)
		}
	}

	return
}


