package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var riyan Customer
	riyan.Name = "riyan"
	riyan.Address = "batang"
	riyan.Age = 23

	fmt.Println(riyan)

	// struct literal
	idin := Customer{
		Address: "pekalongan",
		Age:     18,
		Name:    "idin",
	}
	fmt.Println(idin)

	adit := Customer{"adit", "kendal", 30}
	fmt.Println(adit)
}
