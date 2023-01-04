package main

import "fmt"

type Binatang interface {
	getSuara(string) string
}

type Kucing struct {
	Name string
}

func (kucing Kucing) getSuara(suara string) string {
	return suara + " " + kucing.Name
}

func main() {

	// memanggil singlw data
	riyan := Kucing{"riyan"}
	feri := Kucing{"feri"}
	fmt.Println(riyan.getSuara("aum"))

	// memanggul dengan iterasi
	binatang := []Binatang{riyan, feri}
	for _, bin := range binatang {
		fmt.Println(bin.getSuara("rarr"))
	}
}
