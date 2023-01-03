package main

import "fmt"

// bisa diakses oleh package lain
func SayHello() {
	fmt.Println("hello")
}

// hanya bisa diakses internal package
func sayHi() {
	fmt.Println("hi")
}
