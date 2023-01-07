package main

import "fmt"

func firstGenericFunc[age int64 | float64](myage age) {
	fmt.Println(myage)
}

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func secondGenericFunc[age Number](myage age) {
	fmt.Println(myage)
}

func thirdGenericFunc[age any](myage age) {
	fmt.Println(myage)
}

func bubbleShort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {

	var ageInt int64 = 23
	var ageFloat float64 = 23.4
	var ageString string = "dua puluh tiga"

	firstGenericFunc(ageInt)
	firstGenericFunc(ageFloat)

	secondGenericFunc(ageInt)
	secondGenericFunc(ageFloat)

	thirdGenericFunc(ageInt)
	thirdGenericFunc(ageFloat)
	thirdGenericFunc(ageString)

	sliceInt := []int32{4, 3, 1, 5, 2}
	sliceFloat := []float32{4.1, 3.3, 1.5, 5.3, 2.6, 3.9}
	fmt.Println(bubbleShort(sliceInt))
	fmt.Println(bubbleShort(sliceFloat))
}
