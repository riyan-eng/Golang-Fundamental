package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndo(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{
		City:     "Batang",
		Province: "Jawa tengah",
	}

	fmt.Println(address)

	// argumen harus pointer juga
	ChangeCountryToIndo(&address)
	fmt.Println(address)
}
