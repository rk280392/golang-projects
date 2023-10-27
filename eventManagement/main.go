package main

import (
	"fmt"
	"log"
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

func (d *Date) setDay(day int) error {
	if day <= 0 || day > 31 {
		return fmt.Errorf("Invalid days: %d", day)
	}

	d.Day = day
	return nil
}
func (m *Date) setMonth(month int) error {
	if month <= 0 || month > 12 {
		return fmt.Errorf("Invalid days: %d", month)
	}
	m.Month = month
	return nil
}

func (y *Date) setYear(year int) error {
	if !(year >= 1900 && year <= 2100) {
		return fmt.Errorf("Invalid year: %d", year)
	}
	y.Year = year
	return nil
}
func main() {
	date := Date{}
	err := date.setDay(2)
	if err != nil {
		log.Fatal(err)
	}
	err = date.setMonth(2)
	if err != nil {
		log.Fatal(err)
	}
	err = date.setYear(1992)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date.Day, date.Month, date.Year)
}
