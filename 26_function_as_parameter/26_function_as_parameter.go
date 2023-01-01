package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// type declaration
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello " + nameFiltered)
}

func main() {
	sayHelloWithFilter("riyan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("febri", filter)

	// type declaration
	sayHelloWithFilter2("riyan", spamFilter)
	sayHelloWithFilter2("Anjing", spamFilter)
}
