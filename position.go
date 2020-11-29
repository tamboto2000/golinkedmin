package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Position struct {
	Title        string     `json:"title,omitempty"`
	CompanyName  string     `json:"companyName,omitempty"`
	DateRange    *DateRange `json:"dateRange,omitempty"`
	LocationName string     `json:"locationName,omitempty"`
	Description  string     `json:"description,omitempty"`
	Company      *Company   `json:"company,omitempty"`
}

// compose position from golinkedin.PositionGroup
func composePosition(p *golinkedin.PositionGroup) *Position {
	post := &Position{
		Title:        p.Title,
		CompanyName:  p.CompanyName,
		DateRange:    composeDateRange(p.DateRange),
		LocationName: p.LocationName,
		Description:  p.Description,
		// Company:      composeCompany(p.Company),
	}

	if p.Company != nil {
		post.Company = composeCompany(p.Company)
	}

	return post
}
