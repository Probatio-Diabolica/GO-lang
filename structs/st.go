package main

import (
	"fmt"
)

type wheel struct {
	size int
}

type car struct {
	Name       string
	Sold       int32
	Price      float32
	Type       string
	frontWheel wheel
}

func tire(p car) string {
	if p.Name == "hycross" {
		return "found"
	}
	return "nope"
}

func main() {
	fmt.Print("howdy\n")
	totyta := car{Name: "hycross", Sold: 2100, Price: 900000, frontWheel: wheel{90}}
	fmt.Println(totyta.Sold)
	fmt.Println(tire(totyta))
	totyta.Name = "none"
	fmt.Println(tire(totyta))

	//anon structs
	tata := struct {
		Name string
	}{
		Name: "TATA",
	}

	fmt.Println(tata.Name)
}

// embedded structs
type truck struct {
	car
	tyre int
}
