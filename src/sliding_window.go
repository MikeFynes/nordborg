package main


type SlidingWindow struct{
	times []int
	values []int
	size int
	diff int
	isClosed bool
}

func (w *SlidingWindow) addPoint(time, value int){

	if  len(w.values) < w.size {
		w.times = append(w.times, time)
		w.values = append(w.values, value)
	}
}

func (w *SlidingWindow) wouldClose(time int){

}