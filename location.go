package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Location struct {
	CountryCode    string `json:"countryCode,omitempty"`
	Country        string `json:"country"`
	GeographicArea string `json:"geographicArea,omitempty"`
	City           string `json:"city,omitempty"`
	PostalCode     string `json:"postalCode,omitempty"`
	Description    string `json:"description,omitempty"`
	Headquarter    bool   `json:"headquarter,omitempty"`
	Line2          string `json:"line2,omitempty"`
	Line1          string `json:"line1,omitempty"`
}

func composeLocation(loc *golinkedin.Location) *Location {
	return &Location{
		CountryCode:    loc.Country,
		GeographicArea: loc.GeographicArea,
		City:           loc.City,
		PostalCode:     loc.PostalCode,
		Description:    loc.Description,
		Headquarter:    loc.Headquarter,
		Line1:          loc.Line1,
		Line2:          loc.Line2,
	}
}

// compose Location from golinkedin.Geo
func composeLocationFromGeo(g *golinkedin.Geo) *Location {
	location := &Location{
		City: g.DefaultLocalizedNameWithoutCountryName,
	}

	// extract country
	if g.Country != nil {
		location.Country = g.Country.DefaultLocalizedName
	}

	return location
}
