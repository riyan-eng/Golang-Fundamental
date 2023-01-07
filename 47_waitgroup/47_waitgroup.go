package main

import (
	"fmt"
	"sync"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("getting started goroutines")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("halo")
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("getting started goroutines")

}
