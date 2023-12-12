package vehicles

import "github.com/magalhaes-andre/go-oop/clients"

type Car struct {
	name   string
	price  float64
	plate  string
	driver *clients.Driver
}

func Name(c Car) string {
	return c.name
}
