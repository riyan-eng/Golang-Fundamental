package main

import "fmt"

func main() {
	//
	type NoKTP string
	type Married bool

	var noKtpRiyan NoKTP = "162847126499124"
	var marriedStatus Married = false
	fmt.Println(noKtpRiyan)
	fmt.Println(marriedStatus)
}
