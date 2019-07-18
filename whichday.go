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
	help      int
	day       int
	month     int
	year      int
	cVar      int
	aVar      int
	mVar      int
	j1Var     int
	jVar      int
	weekDay   = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	monthYear = [13]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	nberror   int
	porcelain int
	exit      int
)

// Getting the date from the input arguments
func getInputArguments() {
	argsWithoutProg := os.Args[1:]
	nbArgs := len(argsWithoutProg)
	// Test if arguments were given
	if nbArgs <= 0 {
		return
	}

	// Parse the inputs
	for i := 0; i < nbArgs; i++ {
		theArg := argsWithoutProg[i]
		switch theArg {
		// test if help requested
		case "-h", "--help":
			{
				syntax()
				exit++
				return
			}
		// Get the output format
		case "--porcelain":
			{
				porcelain++
			}
		// Get the day
		case "-d", "--day":
			{
				i++
				value, error := strconv.Atoi(argsWithoutProg[i])
				if error != nil {
					nberror++
				}
				day = value
			}
		// Get the month
		case "-m", "--month":
			{
				i++
				value, error := strconv.Atoi(argsWithoutProg[i])
				if error != nil {
					nberror++
				}
				month = value
			}
		// Get the year
		case "-y", "--year":
			{
				i++
				value, error := strconv.Atoi(argsWithoutProg[i])
				if error != nil {
					nberror++
				}
				year = value
			}
		}
	}
}

func syntax() {
	fmt.Println(`
NAME
    whichday - return the day of the week

SYNOPSIS
    whichday [<options>]

DESCRIPTION
    Determine the day of the week from a given date 
    from 1/11/1582 to 31/12/9999.

OPTIONS
    --porcelain
        Give the output in an easy-to-parse format for scripts.
    -d, --day
        The day from 1 to [28,29,30 or 31]
    -m, --month
        The month from 1 to 12
    -y, -- year
		The year for a date from from 1/11/1582 to 31/12/9999
EXAMPLE
    whichday
        Will interactively ask for the date
    whichday --day 18 --month 7 --year 2019
        will return: 18/07/2019 is Thursday
    whichday --day 18 --month 7 --year 2019 --porcelain
        will return: Thursday
      `)
}

// Getting the date
func getDate() {
	// Get the day if omitted
	if day < 1 {
		value, err := getFromUser("Day:    ")
		if err != nil {
			fmt.Println("Error in day format")
			nberror++
			return
		}
		day = value
	}
	// Get the month if omitted
	if month < 1 {
		value, err := getFromUser("Month:  ")
		if err != nil {
			fmt.Println("Error in month format")
			nberror++
			return
		}
		month = value
	}
	// Get the year if omitted
	if year < 1 {
		value, err := getFromUser("Year:   ")
		if err != nil {
			fmt.Println("Error in year format")
			nberror++
			return
		}
		year = value
	}
}

// Get an input from the console by the user
func getFromUser(message string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return 1, err
	}
	input = strings.TrimRight(input, "\r\n")
	response, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return 1, err
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
}

// Display the result on stdout
func display() {
	if porcelain == 0 {
		fmt.Printf("%s %02d, %04d is %s\n", monthYear[month], day, year, weekDay[jVar])
	} else {
		fmt.Printf("%s\n", weekDay[jVar])
	}
}

func main() {
	getInputArguments()
	if exit > 0 {
		return
	}
	getDate()
	if nberror > 0 {
		fmt.Println("Error was detected. exiting")
		return
	}
	calculate()
	display()
}
