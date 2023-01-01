package main

import "fmt"

func sayHello() {
	fmt.Println("Hello riyan")
}

func main() {
	sayHello()
	sayHello()

	for i := 1; i <= 10; i++ {
		sayHello()
	}
}
