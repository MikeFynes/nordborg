package main

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func readCsv(filePath string)(values []int){
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
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
			values = append(values, i)
		} else {
			log.Fatal(conversionError)
		}
	}

	return
}
