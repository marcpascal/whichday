package main

import (
	"bufio" // Parsing the command line arguments
	"fmt"   // Formatting
	"os"
	"strconv"
	"strings"
)

// Program variable declaration
var (
	help    int
	day     int
	month   int
	year    int
	cVar    int
	aVar    int
	mVar    int
	j1Var   int
	jVar    int
	weekDay [7]string
	nberror int
)

// Init of the week days table
func initVars() {
	weekDay[0] = "Sunday"
	weekDay[1] = "Monday"
	weekDay[2] = "Tuesday"
	weekDay[3] = "Wednesday"
	weekDay[4] = "Thursday"
	weekDay[5] = "Friday"
	weekDay[6] = "Saturday"
	nberror = 0
}

// Getting the date
func getDate() {
	fmt.Println(`Enter the date for which you need the day of the week:
		- day:   from 1 to 31
		- month: from 1 to 12
		- year:  from 1/11/1582 to 31/12/9999`)

	value, err := getFromUser("Day:\t")
	if err != nil {
		fmt.Println("Error in day format")
		nberror++
		return
	}
	day = value
	value, err = getFromUser("Month:\t")
	if err != nil {
		fmt.Println("Error in month format")
		nberror++
		return
	}
	month = value
	value, err = getFromUser("Year:\t")
	if err != nil {
		fmt.Println("Error in year format")
		nberror++
		return
	}
	year = value
}

// Get an input from the console by the user
func getFromUser(message string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	input = strings.TrimRight(input, "\r\n")
	response, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return response, err
}

// Determine the day of the week.
func calculate() {
	cVar = int((14 - month) / 12)
	aVar = year - cVar
	mVar = month + (12 * cVar) - 2
	j1Var = day + aVar +
		int(aVar/4) -
		int(aVar/100) +
		int(aVar/400) +
		int((31*mVar)/12)
	jVar = j1Var % 7
	// fmt.Printf("cVar=%d aVar=%d mVar=%d j1Var=%d jVar=%d\n", cVar, aVar, mVar, j1Var, jVar)
}

// Display the result on stdout
func display() {
	fmt.Printf("\n%02d/%02d/%04d is %s\n", day, month, year, weekDay[jVar])
}

func main() {
	initVars()
	getDate()
	if nberror > 0 {
		fmt.Println("Error was detected. exiting")
		return
	}
	calculate()
	display()
}
