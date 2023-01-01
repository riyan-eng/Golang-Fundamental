package main

import "fmt"

func main() {
	var name = "febriyantoo"

	switch name {
	case "riyan":
		fmt.Println("halo riyan")
	case "febri":
		fmt.Println("halo febri")
	default:
		fmt.Println("kenalan dong")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	// switch tanpa kondisi
	var panjang = len(name)
	switch {
	case panjang > 10:
		fmt.Println("sangat panjang")
	case panjang > 5:
		fmt.Println("terlalu panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
