package golinkedmin

import "github.com/tamboto2000/golinkedin"

type PositionGroup struct {
	DateRange   *DateRange `json:"dateRange,omitempty"`
	CompanyName string     `json:"companyName,omitempty"`
	Positions   []Position `json:"positions,omitempty"`
	Company     *Company   `json:"company,omitempty"`
}

func composePositionGroup(p *golinkedin.PositionGroup) *PositionGroup {
	post := new(PositionGroup)
	if p.DateRange != nil {
		post.DateRange = composeDateRange(p.DateRange)
	}

	post.CompanyName = p.CompanyName
	for _, elm := range p.ProfilePositionInPositionGroup.Elements {
		post.Positions = append(post.Positions, *composePosition(&elm))
	}

	if p.Company != nil {
		post.Company = composeCompany(p.Company)
	}

	return post
}
