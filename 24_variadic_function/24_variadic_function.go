package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total = total + value
	}
	return total
}

func sumAll2(name string, numbers ...int) (string, int) {
	total := 0
	for _, value := range numbers {
		total = total + value
	}
	return name, total
}

func main() {
	total := sumAll(10, 20, 30)
	fmt.Println(total)

	name, totil := sumAll2("riyan", 1, 2, 3)
	fmt.Println(name, totil)

	// slice parameter
	slice := []int{5, 6, 7}
	totol := sumAll(slice...)
	fmt.Println(totol)
}
