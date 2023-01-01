package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("hi", customer.Name)
}

func (customer Customer) sayHuuuu() {
	fmt.Println("huuuuuu", customer.Name)
}

func main() {
	riyan := Customer{
		Address: "batang",
		Age:     23,
		Name:    "riyan",
	}
	riyan.sayHi()

	feri := Customer{Name: "feri"}
	feri.sayHuuuu()
}
