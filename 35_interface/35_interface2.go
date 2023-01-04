package main

import "fmt"

type Binatang interface {
	getSuara(string) string
}

type Kucing struct {
	Name string
}

type Anjing struct {
	Name string
}

func (kucing Kucing) getSuara(suara string) string {
	return suara + " " + kucing.Name
}

func (anjing Anjing) getSuara(suara string) string {
	return suara + " " + anjing.Name
}

func main() {

	// memanggil singlw data
	riyan := Kucing{"riyan"}
	feri := Kucing{"feri"}
	rizal := Kucing{"rizal"}
	fmt.Println(riyan.getSuara("aum"))

	// memanggul dengan iterasi
	binatang := []Binatang{riyan, feri, rizal}
	for _, bin := range binatang {
		fmt.Println(bin.getSuara("rarr"))
	}
}
