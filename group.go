package golinkedmin

import (
	"strconv"
	"strings"
	"time"

	"github.com/tamboto2000/golinkedin"
)

type Group struct {
	ID                  int       `json:"id,omitempty"`
	URL                 string    `json:"url,omitempty"`
	Logo                *Image    `json:"picture,omitempty"`
	BackgroundPicture   *Image    `json:"backgroundPicture,omitempty"`
	GroupName           string    `json:"groupName,omitempty"`
	MemberCount         int       `json:"memberCount,omitempty"`
	Description         string    `json:"description,omitempty"`
	Type                string    `json:"type,omitempty"`
	PostApprovalEnabled bool      `json:"postApprovalEnabled,omitempty"`
	CreatedAt           *Date     `json:"created,omitempty"`
	Rules               string    `json:"rules,omitempty"`
	RelatedGroups       []Group   `json:"relatedGroups,omitempty"`
	Owners              []Profile `json:"owners,omitempty"`
	Admins              []Profile `json:"admins,omitempty"`
}

// compose Group from golinkedin.MiniGroup
func composeMiniGroup(gm *golinkedin.MiniGroup) *Group {
	group := new(Group)

	// extract id
	split := strings.Split(gm.EntityUrn, ":")
	id, _ := strconv.Atoi(split[len(split)-1])
	group.ID = id

	// extract logo
	if gm.Logo.COMLinkedinCommonVectorImage != nil {
		group.Logo = composeImage(gm.Logo.COMLinkedinCommonVectorImage)
	}

	// create group url
	group.URL = "https://www.linkedin.com/groups/" + strconv.Itoa(group.ID)

	// extract group name
	group.GroupName = gm.GroupName

	// extract description
	group.Description = gm.GroupDescription

	return group
}

// compose group from golinkedin.Group
func composeGroup(g *golinkedin.Group) *Group {
	group := new(Group)

	// extract id
	split := strings.Split(g.EntityUrn, ":")
	id, _ := strconv.Atoi(split[len(split)-1])
	group.ID = id

	// create group url
	group.URL = "https://www.linkedin.com/groups/" + strconv.Itoa(group.ID)

	// extract logo
	if g.Logo.COMLinkedinCommonVectorImage != nil {
		group.Logo = composeImage(g.Logo.COMLinkedinCommonVectorImage)
	}

	// extract background picture
	if g.HeroImage != nil {
		if g.HeroImage.COMLinkedinCommonVectorImage != nil {
			group.BackgroundPicture = composeImage(g.HeroImage.COMLinkedinCommonVectorImage)
		}
	}

	// extract group name
	group.GroupName = g.Name.Text

	// extract member count
	group.MemberCount = g.MemberCount

	// extract description
	group.Description = g.Description.Text

	// extract type
	group.Type = g.Type

	// extract post approval enabled
	group.PostApprovalEnabled = g.PostApprovalEnabled

	// extract created at
	group.CreatedAt = unixMilliToDate(int64(g.CreatedAt))

	// extract rules
	group.Rules = g.Rules

	// extract owners
	for _, attr := range g.Owners {
		group.Owners = append(group.Owners, *composeMiniProfile(attr.MiniProfile))
	}

	// extract admins
	for _, attr := range g.Managers {
		group.Admins = append(group.Admins, *composeMiniProfile(attr.MiniProfile))
	}

	return group
}

func unixMilliToDate(t int64) *Date {
	date := time.Unix(t/1000, 0)

	return &Date{
		Year:  date.Year(),
		Month: int(date.Month()),
		Day:   date.Day(),
	}
}
