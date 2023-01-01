package main

import "fmt"

func main() {
	counter := 0
	name := "riyan"
	age := 23

	increment := func() {
		fmt.Println("increment")
		counter++

		name = "feri"
		age := 20
		fmt.Println(age)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println("nama", name)
	fmt.Println("umur", age)
}
