package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Printf(man.Name)
}

func main() {
	riyan := Man{"riyan"}
	riyan.Married()

	fmt.Println(riyan.Name)
}
