package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const logName = "tools-log.txt"
const logFilePermission = 0666

func log(logName string, logFilePermission fs.FileMode) os.File {
	os.Remove(logName)
	logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, logFilePermission)
	if err != nil {
		fmt.Println("There was an error while trying to create log file on path: ", logName)
		fmt.Println("Error: ", err)
	}
	return *logFile
}
func main() {
	log := log(logName, logFilePermission)
	// Full variable declaration
	var author string = "Andre"
	// Inferred variable declaration
	var version = 0.01
	// Shorthand variable declaration
	iterations := 1
	startup(author, version, iterations)
	for {
		input := readInput()
		start := time.Now()
		processInput(input)
		stop := time.Now()
		elapsed := stop.Sub(start)
		log.WriteString(time.Now().Format("Mon 06/01/02 15:04:05 ") + "It took " + elapsed.String() + " to proccess the following input: " + strconv.FormatInt(int64(input), 10) + "\n")
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
	fmt.Println("* -> 5 - Ping list of endpoints, separated by ','")
	fmt.Println("* -> 6 - Ping list of endpoints, contained in a txt file delimited by lines")
	fmt.Scan(&input)
	return input
}

func processInput(input int) {
	var systemOs string = runtime.GOOS
	var dateFormat string = "01-02-2006 15:04:05 Monday"
	var formattedDate string = time.Now().Format(dateFormat)
	var cpuQuantity int = runtime.NumCPU()
	switch input {
	case 0:
		break
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
	case 5:
		var urlsWithCommas string
		fmt.Println("* -> Insert the endpoints you'd like to ping for (separated by comma ','): ")
		fmt.Scan(&urlsWithCommas)
		urlSlice := createSliceFromDelimitedString(urlsWithCommas, ",")
		pingMultipleUrls(urlSlice)
	case 6:
		var filePath string
		fmt.Println("* -> Insert the path for the file containing your list of websites (one website per line): ")
		fmt.Scan(&filePath)
		urls := readUrlsFromFile(filePath)
		pingMultipleUrls(urls)
	default:
		fmt.Println("Unexpected Option")
	}
}

func createSliceFromDelimitedString(line string, delimiter string) []string {
	return strings.Split(line, delimiter)
}

/*
TODO:
Create go application that pings a list of websites inserted through excel or comma string
and checks if they contain the correct protocol available.
*/
func pingMultipleUrls(urls []string) {
	for _, url := range urls {
		ping(url)
	}
}

func ping(url string) {
	start := time.Now()
	response, error := http.Get(url)
	stop := time.Now()
	elapsed := stop.Sub(start)
	if error == nil {
		fmt.Printf("Response status from %s was status %s.\nElapsed time was: %s\n", url, response.Status, elapsed)
	} else {
		fmt.Printf("Request for %s has returned the following error %s\nElapsed time was: %s\n", url, error, elapsed)
	}
}

func readUrlsFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	var urls []string
	if err != nil {
		fmt.Println("There was an error while trying to read the requested file:")
		fmt.Println("File Path: ", filePath)
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("There was an issue while reading the current line: ", line)
			fmt.Println("The error was: ", err)
		}

		urls = append(urls, line)
	}
	file.Close()
	return urls
}
