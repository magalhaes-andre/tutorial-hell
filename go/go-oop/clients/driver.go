package clients

import "github.com/magalhaes-andre/go-oop/vehicles"

type Driver struct {
	name    string
	age     int
	license string
	car     *vehicles.Car
}

func NewDriver(name string, age int, license string) Driver {
	return Driver{name: name, age: age, license: license}
}

func Name(d Driver) string {
	return d.name
}

func Age(d Driver) int {
	return d.age
}

func License(d Driver) string {
	return d.license
}

func GetCar(d Driver) *Car {
	return d.car
}

func HasCar(driver *Driver) bool {
	return driver.car != nil
}
