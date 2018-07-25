package main

import "fmt"

func main() {
	var m = 32
	fmt.Println("hello, world", m)

	var other = Other{"this", 3}
	var another = other.wiblyWobly(5)
	fmt.Println(another)

}
