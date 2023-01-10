package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("riyan-eng", "riyan"))
	fmt.Println(strings.Contains("riyan-eng", "ayub"))
	fmt.Println(strings.Split("riyan eng", " "))
	fmt.Println(strings.Contains("riyan-eng", "riyan"))
	fmt.Println(strings.ToLower("Riyan Eng"))
	fmt.Println(strings.ToUpper("Riyan Eng"))
	fmt.Println(strings.ToTitle("Riyan Eng"))
	fmt.Println(strings.Trim("   riyan eng        ", " "))
	fmt.Println(strings.ReplaceAll("budi budi riyan", "budi", "eng"))
}
