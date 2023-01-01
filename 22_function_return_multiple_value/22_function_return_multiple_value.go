package main

import "fmt"

func getFullName() (string, string) {
	return "riyan", "eng"
}

func main() {
	var firstName, lastName = getFullName()
	fmt.Println(firstName, lastName)

	name1, name2 := getFullName()
	fmt.Println(name1, name2)

	fmt.Println(getFullName())

	// menghiraukan return value
	first, _ := getFullName()
	_, last := getFullName()
	fmt.Println(first, last)
}
