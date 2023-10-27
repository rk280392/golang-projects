package main

import (
	"fmt"
	"log"

	"github.com/rk280392/eventManagement/calander"
)

type Date struct {
	Year  int
	Day   int
	Month int
}

/*

func validateDay(day int) error {
	if day <= 0 || day > 31 {
		return fmt.Errorf("Invalid days: %d", day)
	}
	return nil
}

func validateMonth(month int) error {
	if month <= 0 || month > 12 {
		return fmt.Errorf("Invalid days: %d", month)
	}
	return nil
}

func validateYear(year int) error {
	fmt.Println(year)
	if year >= 1900 && year <= 2100 {
		return fmt.Errorf("Invalid year: %d", year)
	}
	return nil
}

*/

// Instead of using functions like above, we can use methods.

/*
The Date receiver gets a copy of the original struct. Any
updates to the fields of the copy are lost when SetYear exits! So use pointer.
*/

/*
	we added the validation but user can still put invalid values directly through struct

date := Date{}
date.Year = 2019
date.Month = 14
date.Day = 50
fmt.Println(date)

we can move the Date type to another package and make its date fields unexported
*/

func main() {
	date := calander.Date{}
	err := date.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1992)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day, date.Month, date.Year)

	// invalid value is also printed
	date = calander.Date{Year: 0, Month: 20, Day: 40}
	fmt.Println(date)
}
