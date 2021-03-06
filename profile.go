package golinkedmin

import (
	"github.com/tamboto2000/golinkedin"
)

// Profile contains people/user profile
type Profile struct {
	ID                string          `json:"id,omitempty"`
	Username          string          `json:"username,omitempty"`
	ProfilePicture    *Image          `json:"profilePicture,omitempty"`
	BackgroundPicture *Image          `json:"backgroundPicture,omitempty"`
	FirstName         string          `json:"firstName,omitempty"`
	LastName          string          `json:"lastName,omitempty"`
	Headline          string          `json:"headline,omitempty"`
	Location          *Location       `json:"location,omitempty"`
	About             string          `json:"about,omitempty"`
	Premium           bool            `json:"premium,omitempty"`
	Influencer        bool            `json:"influencer,omitempty"`
	Memorialized      bool            `json:"memorialized,omitempty"`
	Experience        []PositionGroup `json:"experience,omitempty"`
	Educations        []Education     `json:"educations,omitempty"`
	Certifications    []Certification `json:"certifications,omitempty"`
	Skills            []Skill         `json:"skills,omitempty"`
	// True if profile composed from golinkedin.Profile
	IsFullProfile bool `json:"isFullProfile,omitempty"`
	// True if profile compose from golinkedin.MiniProfile
	IsMiniProfile bool `json:"isMiniProfile,omitempty"`

	ln *Linkedin
}

// ProfileByName get profile by username
func (ln *Linkedin) ProfileByName(name string) (*Profile, error) {
	prof, err := ln.Linkedin.ProfileByUsername(name)
	if err != nil {
		return nil, parseErrMsg(err.Error())
	}

	composed := composeProfile(&prof.Elements[0])
	composed.ln = ln

	return composed, nil
}

// SetLinkedin set Linkedin client
func (prof *Profile) SetLinkedin(ln *Linkedin) {
	prof.ln = ln
}

// FullProfile get full profile version of Profile
func (prof *Profile) FullProfile() (*Profile, error) {
	raw, err := prof.ln.Linkedin.ProfileByUsername(prof.Username)
	if err != nil {
		return nil, parseErrMsg(err.Error())
	}

	*prof = *composeProfile(&raw.Elements[0])

	return prof, nil
}

// Connections prepare and fetch the first result page of profile connections.
// Use Connections.Next for pagination
func (prof *Profile) Connections() (*Connections, error) {
	raw, err := prof.ln.SearchPeople(
		"",
		&golinkedin.PeopleSearchFilter{
			Network:      []string{golinkedin.Rank1, golinkedin.Rank2, golinkedin.Rank3},
			ConnectionOf: prof.ID,
			ResultType:   golinkedin.ResultPeople,
		},
		&golinkedin.QueryContext{
			SpellCorrectionEnabled: true,
		},
		golinkedin.OriginMemberProfileCannedSearch,
	)

	if err != nil {
		return nil, parseErrMsg(err.Error())
	}

	miniProfs := make([]Profile, 0)
	for _, elm := range raw.Elements {
		if elm.Type == golinkedin.TypeSearchHits {
			for _, innerElm := range elm.Elements {
				miniProfs = append(miniProfs, *composeMiniProfile(innerElm.Image.Attributes[0].MiniProfile))
			}

			break
		}
	}

	conns := &Connections{
		Profiles: miniProfs,
		Start:    raw.Paging.Start,
		Count:    raw.Paging.Count,
		Total:    raw.Paging.Total,
		pnode:    raw,
	}

	return conns, nil
}

// compose Profile from golinkedin.MiniProfile
func composeMiniProfile(m *golinkedin.MiniProfile) *Profile {
	prof := &Profile{
		ID:            extractID(m.EntityUrn),
		Username:      m.PublicIdentifier,
		FirstName:     m.FirstName,
		LastName:      m.LastName,
		Headline:      m.Occupation,
		IsMiniProfile: true,
	}

	// extract background picture
	if m.BackgroundImage != nil {
		prof.BackgroundPicture = composeImage(&m.BackgroundImage.COMLinkedinCommonVectorImage)
	}

	// extract profile picture
	if m.Picture != nil {
		prof.ProfilePicture = composeImage(m.Picture.COMLinkedinCommonVectorImage)
	}

	return prof
}

// compose Profile from golinkedin.Profile
func composeProfile(p *golinkedin.Profile) *Profile {
	prof := &Profile{
		ID:            extractID(p.EntityUrn),
		Username:      p.PublicIdentifier,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		Headline:      p.Headline,
		About:         p.Summary,
		Premium:       p.Premium,
		Influencer:    p.Influencer,
		Memorialized:  p.Memorialized,
		IsFullProfile: true,
	}

	// extract profile picture
	if p.ProfilePicture != nil {
		prof.ProfilePicture = composeImage(p.ProfilePicture.DisplayImageReference.VectorImage)
	}

	// extract background picture
	if p.BackgroundPicture != nil {
		if p.BackgroundPicture.DisplayImageReference != nil {
			prof.BackgroundPicture = &Image{
				URL: p.BackgroundPicture.DisplayImageReference.URL,
			}
		}
	}

	// extract location
	if p.GeoLocation != nil {
		prof.Location = composeLocationFromGeo(&p.GeoLocation.Geo)
	}

	// extract experience
	if p.ProfilePositionGroups != nil {
		for _, post := range p.ProfilePositionGroups.Elements {
			prof.Experience = append(prof.Experience, *composePositionGroup(&post))
		}
	}

	// extract educations
	if p.ProfileEducations != nil {
		for _, edu := range p.ProfileEducations.Elements {
			prof.Educations = append(prof.Educations, *composeEducation(&edu))
		}
	}

	// extract certifications
	if p.ProfileCertifications != nil {
		for _, cert := range p.ProfileCertifications.Elements {
			prof.Certifications = append(prof.Certifications, *composeCertification(&cert))
		}
	}

	// extract skills
	if p.ProfileSkills != nil {
		for _, skill := range p.ProfileSkills.Elements {
			prof.Skills = append(prof.Skills, *composeSkill(&skill))
		}
	}

	return prof
}
