package main

import "fmt"

func main() {
	//
	var names [3]string

	names[0] = "febri"
	names[1] = "riyan"
	names[2] = "eng"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var values = [3]int{
		100, 200, 300,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi = [10]string{}
	fmt.Println(len(lagi))
}
