package main

import "fmt"

func main() {
	presentCar := carHouse("Mersedea-Benz", 2025, 49900)
	fmt.Println(presentCar)
}

type Car struct {
	name  string
	year  int
	price float64
}

func carHouse(name string, year int, price float64) Car {
	return Car{
		name:  name,
		year:  year,
		price: price}
}
