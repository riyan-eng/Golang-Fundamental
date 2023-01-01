package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	factorial := 5 * 4 * 3 * 2 * 1
	fmt.Println(factorial)

	//
	loop := factorialLoop(5)
	fmt.Println(loop)

	//
	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
