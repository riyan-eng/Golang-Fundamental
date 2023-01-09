package main

import (
	"fmt"
	"os"
)

func main() {
	// argument
	args := os.Args
	fmt.Println(args)

	// host sistem operasi
	host, _ := os.Hostname()
	fmt.Println(host)

	// environment
	env := os.Getenv("APP_USERNAME")
	fmt.Println(env)
}
