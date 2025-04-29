package main

import "fmt"

type Address struct {
	Country string
	City    string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	person := Person{
		Name: "Vladimir",
		Age:  25,
		Address: Address{
			Country: "Russia",
			City:    "Moscow",
		},
	}

	fmt.Println(person.Address.City) // Moscow
}
