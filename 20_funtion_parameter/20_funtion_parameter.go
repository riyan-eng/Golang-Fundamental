package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	name := "riyan"
	sayHelloTo(name, "eng")
	sayHelloTo("ahmad", "taali")
}
