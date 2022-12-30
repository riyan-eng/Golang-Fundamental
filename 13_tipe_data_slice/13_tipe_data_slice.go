package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	fmt.Println(months)

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[1] = "Baru"
	// fmt.Println(months)

	//
	var slice2 = months[10:]
	fmt.Println(slice2)
	var slice3 = append(slice2, "riyan")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// make slice
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "febri"
	newSlice[1] = "riyan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	var copySlice = make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// beda slice dan array
	var iniSlice = []int{1, 2, 3, 4, 5}
	var iniArray = [...]int{1, 2, 3, 4, 5}
	fmt.Println(iniSlice)
	fmt.Println(iniArray)
}
