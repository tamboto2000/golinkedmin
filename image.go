package golinkedmin

import "github.com/tamboto2000/golinkedin"

type Image struct {
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	ExpiresAt int64  `json:"expiresAt,omitempty"`
	URL       string `json:"url,omitempty"`
}

func composeImage(v *golinkedin.VectorImage) *Image {
	bestRest := v.Artifacts[len(v.Artifacts)-1]
	img := &Image{
		URL:       v.RootURL + bestRest.FileIdentifyingURLPathSegment,
		Height:    bestRest.Height,
		Width:     bestRest.Width,
		ExpiresAt: bestRest.ExpiresAt,
	}

	return img
}
