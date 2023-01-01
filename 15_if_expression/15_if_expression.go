package main

import "fmt"

func main() {
	var name = "Febriss"

	if name == "Riyan" {
		fmt.Println("Halo Riyan")
	} else if name == "Febri" {
		fmt.Println("Halo Febri")
	} else {
		fmt.Println("Hi kenalan dong")
	}

	//
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
