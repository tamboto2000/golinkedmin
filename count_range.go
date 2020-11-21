package golinkedmin

import "github.com/tamboto2000/golinkedin"

type CountRange struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

func composeCountRange(c *golinkedin.CountRange) *CountRange {
	return &CountRange{
		Start: c.Start,
		End:   c.End,
	}
}
