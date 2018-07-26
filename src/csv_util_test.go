package main

import "testing"

func TestCsvReader(t *testing.T){
	readCsvWithSlidingWindow("Round 1. Software engineering test cases - test2.csv", 5, 25000)
}
