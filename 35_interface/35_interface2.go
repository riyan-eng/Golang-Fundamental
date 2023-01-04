package main

import "fmt"

type Binatang interface {
	getSuara(string) string
}

type Kucing struct {
	Name string
}

func (kucing *Kucing) getSuara(suara string) string {
	return suara + " " + kucing.Name
}

func main() {
	riyan := Kucing{"riyan"}
	fmt.Println(riyan.getSuara("aum"))
}
