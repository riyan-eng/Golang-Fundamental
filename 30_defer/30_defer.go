package main

import "fmt"

func endFunc() {
	fmt.Println("fungsi berhenti")
}

func startFunc() {
	defer endFunc()

	fmt.Println("fungsi berjalan")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	startFunc()
	fmt.Println("halo")
}
