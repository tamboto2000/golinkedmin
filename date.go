package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

func composeDate(d *golinkedin.Date) *Date {
	return &Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}
