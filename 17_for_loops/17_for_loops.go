package main

import "fmt"

func main() {
	var counter = 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// for dengan statement => init dan post
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for range
	var slice = []string{"febri", "riyan", "eng", "na"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	var person = make(map[string]string)
	person["name"] = "riyan"
	person["title"] = "programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
