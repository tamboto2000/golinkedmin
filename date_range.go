package golinkedmin

import "github.com/tamboto2000/golinkedin"

type DateRange struct {
	Start *Date `json:"date,omitempty"`
	End   *Date `json:"end,omitempty"`
}

func composeDateRange(d *golinkedin.DateRange) *DateRange {
	dateR := new(DateRange)
	if d.Start != nil {
		dateR.Start = composeDate(d.Start)
	}

	if d.End != nil {
		dateR.End = composeDate(d.End)
	}

	return dateR
}
