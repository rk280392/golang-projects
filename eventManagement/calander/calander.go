package calander

import "fmt"

type Date struct {
	year  int
	day   int
	month int
}

// Getter method to get individual values

func (d *Date) Day() int {
	return d.day
}

func (m *Date) Month() int {
	return m.month
}
func (y *Date) Year() int {
	return y.year
}

func (d *Date) SetDay(day int) error {
	if day <= 0 || day > 31 {
		return fmt.Errorf("invalid days: %d", day)
	}

	d.day = day
	return nil
}
func (m *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return fmt.Errorf("invalid days: %d", month)
	}
	m.month = month
	return nil
}

func (y *Date) SetYear(year int) error {
	if !(year >= 1900 && year <= 2100) {
		return fmt.Errorf("invalid year: %d", year)
	}
	y.year = year
	return nil
}
