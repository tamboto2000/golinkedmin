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
	Skills    []Skill   `json:"skills,omitmepty"`
	Paging    Paging    `json:"paging,omitempty"`

	ln *Linkedin
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

// SearchPeople search people by keywords and filter
func (ln *Linkedin) SearchPeople(keywords string, filter *golinkedin.PeopleSearchFilter) (*Search, error) {
	return nil, nil
}
