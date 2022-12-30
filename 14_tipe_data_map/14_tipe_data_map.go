package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "riyan",
		"address": "batang",
	}
	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//
	var books = make(map[string]string)
	books["title"] = "go fun"
	books["author"] = "riyan"
	books["say"] = "hello"

	fmt.Println(books)
	fmt.Println(len(books))

	delete(books, "say")
	fmt.Println(books)
	fmt.Println(len(books))
}
