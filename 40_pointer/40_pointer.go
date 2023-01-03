package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// menggunakan &
	address1 := Address{
		City:     "Batang",
		Province: "Jawa tengah",
		Country:  "Indonesia",
	}

	address2 := &address1
	address2.City = "Pekalongan"
	fmt.Println(address1)
	fmt.Println(address2)

	// menggunakan *
	*address2 = Address{
		City:     "Pemalang",
		Province: "Jawa tengah",
		Country:  "Indonesia",
	}
	fmt.Println(address1)
	fmt.Println(address2)

	// menggunakan new
	var address3 = new(Address)
	address3.City = "Bandar"
	fmt.Println(address3)
}
