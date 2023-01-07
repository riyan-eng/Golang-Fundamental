package main

import (
	"fmt"
	"time"
)

func sender(c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func receiver(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan string)

	go sender(c)
	go receiver(c)

	fmt.Scanln()
}
