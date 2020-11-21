package golinkedmin

import (
	"strconv"
	"strings"

	"github.com/tamboto2000/golinkedin"
)

type Company struct {
	ID                   int      `json:"id,omitempty"`
	UniversalName        string   `json:"universalName,omitempty"`
	URL                  string   `json:"url,omitempty"`
	Logo                 *Image   `json:"logo,omitempty"`
	BackgroundCoverImage *Image   `json:"backgroundCoverImage,omitempty"`
	CompanyName          string   `json:"companyName,omitempty"`
	Industries           []string `json:"industry,omitempty"`
	FollowerCount        int      `json:"followerCount,omitempty"`
	Tagline              string   `json:"tagline,omitempty"`
	Description          string   `json:"description,omitempty"`
	// StaffCountRange is staff count range based on company's admin input
	StaffCountRange *CountRange `json:"staffCountRange,omitempty"`
	// StaffCount is staff count based on people on linkedin worked in the company
	StaffCount          int        `json:"staffCount,omitempty"`
	Headquarter         *Location  `json:"headquarter,omitempty"`
	CompanyType         string     `json:"companyType,omitempty"`
	FoundedOn           *Date      `json:"founded,omitempty"`
	Specialities        []string   `json:"specialities,omitempty"`
	Locations           []Location `json:"locations,omitempty"`
	AffiliatedCompanies []Company  `json:"affiliatedCompanies,omitempty"`
	FeaturedGroups      []Group    `json:"featuredGroups,omitempty"`
}

// compose company from golinkedin.Company
func composeCompany(c *golinkedin.Company) *Company {
	comp := &Company{
		UniversalName: c.UniversalName,
		URL:           c.URL,
		CompanyName:   c.Name,
		Tagline:       c.Tagline,
		Description:   c.Description,
		StaffCount:    c.StaffCount,
	}

	// extract id
	split := strings.Split(c.EntityUrn, ":")
	id, _ := strconv.Atoi(split[len(split)-1])
	comp.ID = id

	// extract logo
	// from Company.Logo.VectorImage
	if c.Logo.VectorImage != nil {
		comp.Logo = composeImage(c.Logo.VectorImage)
	}

	// extract background cover image
	if c.BackgroundCoverImage != nil {
		if c.BackgroundCoverImage.Image.COMLinkedinCommonVectorImage != nil {
			comp.BackgroundCoverImage = composeImage(c.BackgroundCoverImage.Image.COMLinkedinCommonVectorImage)
		}
	}

	// from Company.Logo.Image.COMLinkedinCommonVectorImage
	if c.Logo.Image != nil {
		if c.Logo.Image.COMLinkedinCommonVectorImage != nil {
			comp.Logo = composeImage(c.Logo.COMLinkedinCommonVectorImage)
		}
	}

	// extract industries
	// from golinkedin.Company.Industry
	if c.Industry != nil {
		for _, val := range c.Industry {
			comp.Industries = append(comp.Industries, val.Name)
		}
	}

	// from golinkedin.Company.CompanyIndustries
	if c.CompanyIndustries != nil {
		for _, val := range c.CompanyIndustries {
			comp.Industries = append(comp.Industries, val.LocalizedName)
		}
	}

	// extract follower count
	if c.FollowingInfo != nil {
		comp.FollowerCount = c.FollowingInfo.FollowerCount
	}

	// extract staff count range
	// from golinkedin.Company.EmployeeCountRange.Start
	if c.EmployeeCountRange != nil {
		comp.StaffCountRange = composeCountRange(c.EmployeeCountRange)
	}

	// from golinkedin.Company.StaffCountRange
	if c.StaffCountRange != nil {
		comp.StaffCountRange = composeCountRange(c.StaffCountRange)
	}

	// extract headquarter
	if c.Headquarter != nil {
		comp.Headquarter = &Location{
			Country:    c.Headquarter.Country,
			City:       c.Headquarter.City,
			PostalCode: c.Headquarter.PostalCode,
			Line1:      c.Headquarter.Line1,
			Line2:      c.Headquarter.Line2,
		}
	}

	// extract company type
	if c.CompanyType != nil {
		comp.CompanyType = c.CompanyType.LocalizedName
	}

	// extract founded date
	if c.FoundedOn != nil {
		comp.FoundedOn = composeDate(c.FoundedOn)
	}

	// extract specialities
	if c.Specialities != nil {
		comp.Specialities = append(comp.Specialities, c.Specialities...)
	}

	// extract locations
	if c.ConfirmedLocations != nil {
		for _, loc := range c.ConfirmedLocations {
			comp.Locations = append(comp.Locations, Location{
				Country:     loc.Country,
				City:        loc.City,
				PostalCode:  loc.PostalCode,
				Description: loc.Description,
				Headquarter: loc.Headquarter,
				Line1:       loc.Line1,
				Line2:       loc.Line2,
			})
		}
	}

	// extract affiliated companies
	if c.AffiliatedCompaniesResolutionResults != nil {
		for _, val := range c.AffiliatedCompaniesResolutionResults {
			comp.AffiliatedCompanies = append(comp.AffiliatedCompanies, *composeCompany(&val))
		}
	}

	// extract featured groups
	if c.GroupsResolutionResults != nil {
		for _, val := range c.GroupsResolutionResults {
			comp.FeaturedGroups = append(comp.FeaturedGroups, *composeGroup(&val))
		}
	}

	return comp
}
