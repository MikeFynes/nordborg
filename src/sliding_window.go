package main

import "sort"

type SlidingWindow struct{
	times []int // milliseconds since start
	values []int
	size int // number of entries
	timeDiffMax int // max difference between first entry and latest
}

func (w *SlidingWindow) addDelay(time, value int){

	// check time vs first entry in window, if greater than max then keep sliding until it is not
	for len(w.times) > 0 && time - w.times[0] > w.timeDiffMax {
		slideWindow(w)
	}

	if len(w.times) == w.size {
		slideWindow(w)
	}

	w.times = append(w.times, time)
	w.values = append(w.values, value)
}

func slideWindow(w *SlidingWindow) {
	w.times = w.times[1:]
	w.values = w.values[1:]
}

//If only one element available in the sliding window the answer is -1.
//If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
//If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ]
func (w *SlidingWindow) getMedian()(median int){
	var values = make([]int, len(w.values))
	copy(values, w.values)
	sort.Ints(values)
	switch{
	case len(values) == 1 : median = -1
	case len(values) % 2 == 0 :
		median = (values[len(values) / 2 -1] + values[len(values) / 2]) / 2
	case len(values) % 2 != 0 :
		var itemOne = (len(values) + 1) / 2 -1
		median = values[itemOne]
	}

	return
}