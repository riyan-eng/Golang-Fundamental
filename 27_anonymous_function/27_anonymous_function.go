package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", blacklist)
	registerUser("riyan", blacklist)

	registerUser("babi", func(name string) bool {
		return name == "babi"
	})
	registerUser("febri", func(name string) bool {
		return name == "babi"
	})
}
