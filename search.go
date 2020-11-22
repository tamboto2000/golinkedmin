package golinkedmin

import (
	"github.com/tamboto2000/golinkedin"
)

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
	Profiles            []Profile                      `json:"profiles,omitempty"`
	Companies           []Company                      `json:"companies,omitempty"`
	Schools             []School                       `json:"schools,omitempty"`
	Groups              []Group                        `json:"groups,omitempty"`
	Skills              []Skill                        `json:"skills,omitempty"`
	Type                string                         `json:"type,omitempty"`
	Paging              Paging                         `json:"paging,omitempty"`
	Keywords            string                         `json:"keywords,omitempty"`
	ProfileSearchFilter *golinkedin.PeopleSearchFilter `json:"profileSearchFilter,omitempty"`

	err error
	ln  *Linkedin
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

	if filter == nil {
		filter = golinkedin.DefaultSearchPeopleFilter
	}

	search := &Search{
		Type:                SearchProfile,
		ProfileSearchFilter: filter,
		Keywords:            keywords,
		ln:                  ln,
	}

	search.Paging = composePaging(res.Paging)
	for _, elem := range res.Elements {
		if elem.Type == "SEARCH_HITS" {
			for _, p := range elem.Elements {
				prof := *composeMiniProfile(p.Image.Attributes[0].MiniProfile)
				prof.ln = ln
				search.Profiles = append(search.Profiles, prof)
			}

			break
		}
	}

	return search, nil
}

// Next execute cursor
func (srch *Search) Next() bool {
	if srch.Type == SearchProfile {
		return srch.nextProfile()
	}

	return false
}

// Error return error from cursoring operation
func (srch *Search) Error() error {
	return srch.err
}

// cursor for profile search
func (srch *Search) nextProfile() bool {
	node := &golinkedin.PeopleNode{
		Keywords:     srch.Keywords,
		Filters:      srch.ProfileSearchFilter,
		QueryContext: golinkedin.DefaultSearchPeopleQueryContext,
		Paging: golinkedin.Paging{
			Start: srch.Paging.Start,
			Count: srch.Paging.Count,
			Total: srch.Paging.Total,
		},
	}

	node.SetLinkedin(srch.ln.Linkedin)
	if node.Next() {
		profs := make([]Profile, 0)
		for _, elem := range node.Elements {
			if elem.Type == "SEARCH_HITS" {
				for _, p := range elem.Elements {
					prof := *composeMiniProfile(p.Image.Attributes[0].MiniProfile)
					prof.ln = srch.ln
					profs = append(profs, prof)
				}

				break
			}
		}

		srch.Profiles = profs
		srch.Paging = composePaging(node.Paging)

		return true
	}

	if err := node.Error(); err != nil {
		srch.err = err
	}

	return false
}
