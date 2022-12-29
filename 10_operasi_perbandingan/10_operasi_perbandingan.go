package main

import "fmt"

func main() {
	var name1 = "riyan"
	var name2 = "riyan"
	var name3 = "febri"

	var result1 bool = name1 == name2
	fmt.Println(result1)

	var result2 bool = name1 != name3
	fmt.Println(result2)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
