package golinkedmin

// Search contains search result
type Search struct {
	Profiles  []Profile `json:"profiles,omitempty"`
	Companies []Company `json:"companies,omitempty"`
	Schools   []School  `json:"schools,omitempty"`
	Groups    []Group   `json:"groups,omitempty"`
	Paging    Paging    `json:"paging,omitempty"`
	Skills    []Skill   `json:"skills,omitmepty"`

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
