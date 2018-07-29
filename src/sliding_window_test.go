package main

import (
	"testing"
)

func TestSlidingWindowMedianCalculation(t *testing.T){
	var expected = []int{-1, 101, 101, 102, 110, 115}

	var time = []int{100, 200, 300, 400, 500, 600}
	var input = []int{100, 102, 101, 110, 120, 115}

	var window = SlidingWindow{}
	window.size = 3
	window.timeDiffMax = 25000
	for i := 0; i < len(input); i++ {
		window.addDelay(time[i], input[i])
		var median = window.getMedian()
		if median != expected[i] {
			t.Errorf("Median was incorrect, got: %d, want: %d.", median, expected[i])

		}
	}
}

func TestSlidingWindowTimeDiffMaxSlide(t *testing.T) {
	var expected = []int{-1, 101, -1, 105, 115, 115}

	var time = []int{100, 150, 25000, 25200, 25300, 25400}
	var input = []int{100, 102, 101, 110, 120, 115}
	var window = SlidingWindow{}
	window.size = 3
	window.timeDiffMax = 200
	for i := 0; i < len(input); i++ {
		window.addDelay(time[i], input[i])
		var median = window.getMedian()
		if median != expected[i] {
			t.Errorf("Median was incorrect, got: %d, want: %d.", median, expected[i])
		}
	}

}
