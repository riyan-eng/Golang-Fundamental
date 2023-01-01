package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "halo bro"
	} else {
		return "halo " + name
	}
}

func main() {
	result := getHello("riyan")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
