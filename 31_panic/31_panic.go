package main

import "fmt"

func endFunc() {
	fmt.Println("fungsi berhenti")
}

func startFunc(error bool) {
	defer endFunc()
	fmt.Println("fungsi berjalan")
	if error {
		panic("fungsi eror")
	}
	fmt.Println("riyan")
}

func main() {
	startFunc(true)
	fmt.Println("halo")
}
