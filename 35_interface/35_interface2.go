package main

import "fmt"

type Binatang interface {
	getSuara(string) string
}

type Kucing struct {
	Name string
}

func (kucing Kucing) getSuara() string {
	return "meong " + kucing.Name
}

func main() {
	riyan := Kucing{"riyan"}
	fmt.Println(riyan.getSuara())
}
