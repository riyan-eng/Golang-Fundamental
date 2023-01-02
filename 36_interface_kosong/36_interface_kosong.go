package main

import "fmt"

func ups(value int) interface{} {
	if value == 1 {
		return 1
	} else if value == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = ups(3)
	fmt.Println(data)
}
