package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Education struct {
	Activities   string     `json:"activities,omitempty"`
	School       *School    `json:"school,omitempty"`
	Description  string     `json:"description,omitempty"`
	DegreeName   string     `json:"degreeName,omitempty"`
	SchoolName   string     `json:"schoolName,omitempty"`
	FieldOfStudy string     `json:"fieldOfStudy,omitempty"`
	DateRange    *DateRange `json:"dateRange,omitempty"`
}

func composeEducation(e *golinkedin.Education) *Education {
	edu := &Education{
		Activities:   e.Activities,
		School:       composeSchool(e.School),
		Description:  e.Description,
		DegreeName:   e.DegreeName,
		SchoolName:   e.SchoolName,
		FieldOfStudy: e.FieldOfStudy,
		DateRange:    composeDateRange(e.DateRange),
	}

	return edu
}
