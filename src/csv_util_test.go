package main

import "testing"

func TestCsvReaderDoc1(t *testing.T){
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test2.csv", 5, 25000)
}

func TestCsvReaderDoc2(t *testing.T){
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test3.csv", 5, 25000)
}

func TestCsvReaderDoc3(t *testing.T){
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test4.csv", 5, 25000)
}
