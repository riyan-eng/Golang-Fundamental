package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	fmt.Println(boolean)
	fmt.Println(err)

	integer, err := strconv.ParseInt("1000000", 10, 64)
	fmt.Println(integer)
	fmt.Println(err)

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("1000000")
	fmt.Println(valueInt)

	valueStr := strconv.Itoa(10000)
	fmt.Println(valueStr)
}
