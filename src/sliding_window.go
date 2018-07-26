package main

import "sort"

type SlidingWindow struct{
	times []int
	values []int
	size int
	diff int
}

func (w *SlidingWindow) addPoint(time, value int){
	// Slide window
	if len(w.times) == w.size {
		var reversedTimes = reverseIntArray(w.times)
		var reversedValues = reverseIntArray(w.values)
		w.times = reverseIntArray(reversedTimes[:w.size-1])
		w.values = reverseIntArray(reversedValues[:w.size-1])
	}

	w.times = append(w.times, time)
	w.values = append(w.values, value)
}

//If only one element available in the sliding window the answer is -1.
//If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
//If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ]
func (w *SlidingWindow) medianOfWindow()(median int){
	var values = make([]int, len(w.values))
	copy(values, w.values)
	sort.Ints(values)
	switch{
	case len(values) == 1 : median = -1
	case len(values) % 2 == 0 :
		var itemOne = len(values) / 2 -1
		var itemTwo = (len(values) / 2) +1 -1
		var interimOne = values[itemOne]
		var interimTwo = values[itemTwo]
		median = (interimOne + interimTwo) / 2
	case len(values) % 2 != 0 :
		var itemOne = (len(values) + 1) / 2 -1
		median = values[itemOne]
	}

	return
}

func reverseIntArray(s []int)(reversed []int){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	reversed = s

	return
}