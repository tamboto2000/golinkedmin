package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Certification struct {
	DateRange       *DateRange `json:"dateRange,omitempty"`
	URL             string     `json:"url,omitempty"`
	Authority       string     `json:"authority,omitempty"`
	CertificateName string     `json:"certificateName,omitempty"`
	LicenseNumber   string     `json:"licenseNumber,omitempty"`
	Company         *Company   `json:"company,omitempty"`
	Source          string     `json:"source,omitempty"`
}

func composeCertification(c *golinkedin.Certification) *Certification {
	cert := &Certification{
		DateRange:       composeDateRange(&c.DateRange),
		URL:             c.URL,
		Authority:       c.Authority,
		CertificateName: c.Name,
		LicenseNumber:   c.LicenseNumber,
		Company:         composeCompany(&c.Company),
		Source:          c.DisplaySource,
	}

	return cert
}
