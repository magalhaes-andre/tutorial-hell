package main

import (
	"log"

	"github.com/magalhaes-andre/go-oop/clients"
)

// TODO: finish rewriting the classes
func main() {
	adam := clients.NewDriver("Adam")

	benson := models.Driver{
		Name:    "Benson",
		Age:     55,
		License: "18CDIK67",
	}

	toyota := models.Car{
		Name:  "Toyota",
		Price: 5420.23,
		Plate: "F17F22",
	}

	deliverCarToDealership(&toyota, &adam)
	deliverCarToDriver(&toyota, &adam)
	deliverCarToDriver(&toyota, &adam)
	deliverCarToDealership(&toyota, &adam)
	transferCar(&adam, &benson)
	deliverCarToDriver(&toyota, &benson)
	transferCar(&benson, &adam)
}

func deliverCarToDealership(car *models.Car, driver *models.Driver) {
	if hasCar(driver) {
		log.Printf("Account %s contains the following car assigned: %s", driver.Name, car.Name)
		if driver.Car == car {
			driver.Car = nil
			log.Printf("Car %s has been removed from %s account. Current Driver: %v", car.name, driver.name, driver)
		}
		return
	}

	log.Printf("Account %s does not have any car assigned: %v", driver.name, driver)
}

func deliverCarToDriver(car *Car, driver *Driver) {
	if hasCar(driver) {
		log.Printf("Account %s already has a car: %s", driver.name, car.name)
		log.Printf("Please deliver before assigning a car.")
		return
	}

	driver.car = car
	car.driver = driver
	log.Printf("The car %s has been assigned to the driver %s, their references are now: \n Driver: %v, \n Car: %v", car.name, driver.name, driver, car)
}

func transferCar(owner *Driver, newOwner *Driver) {
	if hasCar(owner) && !hasCar(newOwner) {
		log.Printf("The current owner of car %s will be transferring his car to the new owner: %s\n Owner: %v,\n New Owner: %v", owner.car.name, newOwner.name, owner, newOwner)
		newOwner.car = owner.car
		owner.car = nil
		log.Printf("Transference has been finished: \n Current Owner: %v, \n Old Owner: %v", newOwner, owner)
		return
	}

	log.Printf("Transference is unable to happen at the moment: \n	Owner: %v, \n	New Owner: %v", owner, newOwner)
}

func hasCar(driver *Driver) bool {
	return driver.car != nil
}
