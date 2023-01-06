package main

import "fmt"

func compute(value int) {
	for i := 0; i < value; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("getting started goroutines")
	go compute(5)
	go compute(5)

	fmt.Scanln()
}
