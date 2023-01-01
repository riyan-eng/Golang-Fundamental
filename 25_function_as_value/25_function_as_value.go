package main

import "fmt"

func getGoodBye(name string) string {
	return "bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("riyan")
	fmt.Println(result)

	fmt.Println(getGoodBye("riyan"))
}
