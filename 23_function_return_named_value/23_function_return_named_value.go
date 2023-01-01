package main

import "fmt"

func getFullName() (firstName string, lastName string) {
	firstName = "riyan"
	lastName = "eng"
	return
}

func main() {
	a, b := getFullName()
	fmt.Println(a, b)
}
