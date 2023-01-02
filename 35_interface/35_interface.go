package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasname HasName) {
	fmt.Println("halo", hasname.GetName())
}

type Person struct {
	Nama string
}

func (person Person) GetName() string {
	return person.Nama
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	riyan := Person{Nama: "riyan"}
	anjing := Animal{Name: "anjing"}

	sayHello(riyan)
	sayHello(anjing)
}
