package golinkedmin

import (
	"strconv"

	"github.com/tamboto2000/golinkedin"
)

type School struct {
	ID                   int      `json:"id,omitempty"`
	UniversalName        string   `json:"universalName,omitempty"`
	URL                  string   `json:"url,omitempty"`
	Logo                 *Image   `json:"logo,omitempty"`
	BackgroundCoverImage *Image   `json:"backgroundCoverImage,omitempty"`
	Industries           []string `json:"industries,omitempty"`
	FollowerCount        int      `json:"followerCount,omitempty"`
	Tagline              string   `json:"tagline,omitempty"`
	Description          string   `json:"description,omitempty"`
	// StaffCountRange is staff count range based on company's admin input
	StaffCountRange *CountRange `json:"staffCountRange,omitempty"`
	// StaffCount is staff count based on people on linkedin worked in the company
	StaffCount        int        `json:"staffCount,omitempty"`
	Headquarter       *Location  `json:"headquarter,omitempty"`
	CompanyType       string     `json:"companyType,omitempty"`
	FoundedOn         *Date      `json:"founded,omitempty"`
	Locations         []Location `json:"locations,omitempty"`
	AffiliatedSchools []School   `json:"affiliatedSchools,omitempty"`
	FeaturedGroups    []Group    `json:"featuredGroups,omitempty"`
}

func composeMiniSchool(sch *golinkedin.MiniSchool) *School {
	// extract school id
	idStr := extractID(sch.EntityUrn)
	id, _ := strconv.Atoi(idStr)

	school := &School{
		ID:  id,
		URL: "https://www.linkedin.com/school/" + strconv.Itoa(id),
	}

	// extract logo
	if sch.Logo.VectorImage != nil {
		school.Logo = composeImage(sch.Logo.VectorImage)
	} else if sch.Logo.COMLinkedinCommonVectorImage != nil {
		school.Logo = composeImage(sch.Logo.COMLinkedinCommonVectorImage)
	}

	return school
}

func composeSchool(sch *golinkedin.School) *School {
	idStr := extractID(sch.EntityUrn)
	id, _ := strconv.Atoi(idStr)
	school := &School{
		ID:            id,
		UniversalName: sch.UniversalName,
		URL:           sch.URL,
		Tagline:       sch.Tagline,
		Description:   sch.Description,
		StaffCount:    sch.StaffCount,
	}

	// extract logo
	if sch.Logo != nil {
		if sch.Logo.VectorImage != nil {
			school.Logo = composeImage(sch.Logo.VectorImage)
		}

		if sch.Logo.COMLinkedinCommonVectorImage != nil {
			school.Logo = composeImage(sch.Logo.COMLinkedinCommonVectorImage)
		}
	}

	// extract background cover image
	if sch.BackgroundCoverImage != nil {
		school.BackgroundCoverImage = composeImage(sch.BackgroundCoverImage.Image.COMLinkedinCommonVectorImage)
	}

	// extract industries
	if sch.CompanyIndustries != nil {
		for _, ind := range sch.CompanyIndustries {
			school.Industries = append(school.Industries, ind.LocalizedName)
		}
	}

	// extract follower count
	if sch.FollowingInfo != nil {
		school.FollowerCount = sch.FollowingInfo.FollowerCount
	}

	// extract staff count range
	if sch.StaffCountRange != nil {
		school.StaffCountRange = composeCountRange(sch.StaffCountRange)
	}

	// extract headquarter
	if sch.Headquarter != nil {
		school.Headquarter = composeLocation(sch.Headquarter)
	}

	// extract company type
	if sch.CompanyType != nil {
		school.CompanyType = sch.CompanyType.LocalizedName
	}

	// extract founded
	if sch.FoundedOn != nil {
		school.FoundedOn = composeDate(sch.FoundedOn)
	}

	// extract locations
	if sch.ConfirmedLocations != nil {
		for _, loc := range sch.ConfirmedLocations {
			school.Locations = append(school.Locations, *composeLocation(&loc))
		}
	}

	// extract affiliated schools
	if sch.AffiliatedCompaniesResolutionResults != nil {
		for _, sch := range sch.AffiliatedCompaniesResolutionResults {
			school.AffiliatedSchools = append(school.AffiliatedSchools, *composeSchool(&sch))
		}
	}

	// extract affiliated group
	if sch.GroupsResolutionResults != nil {
		for _, gr := range sch.GroupsResolutionResults {
			school.FeaturedGroups = append(school.FeaturedGroups, *composeGroup(&gr))
		}
	}

	return school
}
