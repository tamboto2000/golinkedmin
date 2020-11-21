package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Skill struct {
	Name      string    `json:"name,omitempty"`
	Endorsers []Profile `json:"endorsers,omitempty"`
}

func composeSkill(s *golinkedin.Skill) *Skill {
	return &Skill{
		Name: s.Name,
	}
}
