package main

import "fmt"

type Car struct {
	name   string
	price  float64
	plate  string
	driver *Driver
}

type Driver struct {
	name    string
	age     int
	license string
	car     *Car
}

func main() {
	driver := Driver{
		name:    "Adam",
		age:     23,
		license: "17ABUJ56",
	}

	var anotherDriver *Driver
	anotherDriver = new(Driver)
	anotherDriver.name = "Benson"
	anotherDriver.age = 35
	anotherDriver.license = "18CDIK67"

	car := Car{
		name:   "Toyota",
		price:  5420.23,
		plate:  "F17F22",
		driver: &driver,
	}

	fmt.Println("Car: ", car)
	fmt.Println("Driver: ", driver)

	driver.car = &car

	fmt.Println("Car: ", car)
	fmt.Println("Driver: ", driver)
	fmt.Println("Another Driver: ", anotherDriver)
}
