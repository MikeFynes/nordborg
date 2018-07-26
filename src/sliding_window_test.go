package main

import (
	"testing"
)

func TestSlidingWindow(t *testing.T){
	var expected = []int{-1, 101, 101, 102, 110, 115}

	var time = []int{0, 5, 10, 15, 20, 25}
	var input = []int{100, 102, 101, 110, 120, 115}

	var window = SlidingWindow{}
	window.size = 3
	for i := 0; i < len(input); i++ {
		window.addPoint(time[i], input[i])
		var median = window.medianOfWindow()
		if median != expected[i] {
			t.Errorf("Median was incorrect, got: %d, want: %d.", median, expected[i])

		}
	}
}
