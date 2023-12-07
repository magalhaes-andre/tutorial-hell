package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	// Full variable declaration
	var author string = "Andre"
	// Inferred variable declaration
	var version = 0.01
	// Shorthand variable declaration
	iterations := 1
	startup(author, version, iterations)
	for {
		input := readInput()
		processInput(input)
	}
}

func startup(author string, version float64, iterations int) {
	fmt.Println("Program Author:", author)
	fmt.Println("Program Version:", version)
	fmt.Println("Code Iterations:", iterations)
}

func readInput() int {
	var input int
	// %d represents int values for fmt.Scanf("%d", intVariable)
	// &usage on variable due to the need of using the exact pointer to attribute value to the variable. TODO: research on pointers in GO
	// If you wish, there is also fmt.Scan only that does not need the specification on which type is being expected. TODO: research on fmt.Scan usages advantages and demises
	//fmt.Scan(&input)
	fmt.Println("What would you like to do?")
	fmt.Println("* -> 1 - Fetch Current System OS")
	fmt.Println("* -> 2 - Fetch Current System Date and Time")
	fmt.Println("* -> 3 - Fetch Number of CPUs in the System")
	fmt.Println("* -> 4 - Ping endpoint of your choice")
	fmt.Scan(&input)
	return input
}

func processInput(input int) {
	var systemOs string = runtime.GOOS
	var dateFormat string = "01-02-2006 15:04:05 Monday"
	var formattedDate string = time.Now().Format(dateFormat)
	var cpuQuantity int = runtime.NumCPU()
	switch input {
	case 1:
		fmt.Println("* -> Current Running System:", systemOs)
	case 2:
		fmt.Println("* -> Current Date:", formattedDate)
	case 3:
		fmt.Println("* -> Quantity of CPUs in the System:", cpuQuantity)
	case 4:
		var url string
		fmt.Println("* -> Insert the endpoint you'd like to ping for: ")
		fmt.Scan(&url)
		ping(url)
	default:
		fmt.Println("Unexpected Option")
	}
}

func ping(url string) {
	response, error := http.Get(url)
	if error == nil {
		fmt.Printf("Response status from %s with status %s.", url, response.Status)
	} else {
		fmt.Printf("Request for %s has returned the following error %s", url, error)
	}
}
