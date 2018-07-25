package main

type Other struct {
	name string
	wibble int
}

func (o *Other) wiblyWobly (woble int) (another Other) {
		another.wibble = o.wibble * woble
		another.name = o.name + " wobbled!"
		return
}