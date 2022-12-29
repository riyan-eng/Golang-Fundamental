package main

import "fmt"

func main() {
	//
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// augmented assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// unary operator
	i++
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
