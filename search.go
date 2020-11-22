package golinkedmin

import "github.com/tamboto2000/golinkedin"

// Search types
const (
	SearchProfile = "PROFILE"
	SearchCompany = "COMPANY"
	SearchSchool  = "SCHOOL"
	SearchGroup   = "GROUP"
	SearchSkill   = "SKILL"
)

// Search contains search result
type Search struct {
	Profiles  []Profile `json:"profiles,omitempty"`
	Companies []Company `json:"companies,omitempty"`
	Schools   []School  `json:"schools,omitempty"`
	Groups    []Group   `json:"groups,omitempty"`
	Skills    []Skill   `json:"skills,omitempty"`
	Type      string    `json:"type,omitempty"`
	Paging    Paging    `json:"paging,omitempty"`

	ln *Linkedin
}

// compose Paging from golinkedin.Paging
func composePaging(p golinkedin.Paging) Paging {
	return Paging{
		Start: p.Start,
		Count: p.Count,
		Total: p.Total,
	}
}

// Paging contains result cursor info.
// We recommend to NOT edit Start value after first result, Count is safe to edit
type Paging struct {
	Start int `json:"start,omitempty"`
	Count int `json:"count,omitempty"`
	Total int `json:"total,omitempty"`
}

// SetLinkedin set Linkedin client
func (srch *Search) SetLinkedin(ln *Linkedin) {
	srch.ln = ln
}

// SearchProfile search people by keywords and filter
func (ln *Linkedin) SearchProfile(keywords string, filter *golinkedin.PeopleSearchFilter) (*Search, error) {
	res, err := ln.Linkedin.SearchPeople(keywords, filter)
	if err != nil {
		return nil, err
	}

	search := &Search{
		Type: SearchProfile,
		ln:   ln,
	}
	search.Paging = composePaging(res.Paging)
	for _, p := range res.Elements[0].Elements {
		search.Profiles = append(search.Profiles, *composeMiniProfile(p.Image.Attributes[0].MiniProfile))
	}

	return search, nil
}
