package calander

import "fmt"

type Date struct {
	Year  int
	Day   int
	Month int
}

func (d *Date) SetDay(day int) error {
	if day <= 0 || day > 31 {
		return fmt.Errorf("invalid days: %d", day)
	}

	d.Day = day
	return nil
}
func (m *Date) SetMonth(month int) error {
	if month <= 0 || month > 12 {
		return fmt.Errorf("invalid days: %d", month)
	}
	m.Month = month
	return nil
}

func (y *Date) SetYear(year int) error {
	if !(year >= 1900 && year <= 2100) {
		return fmt.Errorf("invalid year: %d", year)
	}
	y.Year = year
	return nil
}
