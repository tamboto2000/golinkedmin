package golinkedmin

import (
	"strings"

	"github.com/tamboto2000/golinkedin"
)

type Linkedin struct {
	*golinkedin.Linkedin
}

func New() *Linkedin {
	return &Linkedin{Linkedin: golinkedin.New()}
}

func (ln *Linkedin) SetCookieStr(c string) {
	ln.Linkedin.SetCookieStr(c)
}

// generic id extractor
func extractID(str string) string {
	split := strings.Split(str, ":")
	return split[len(split)-1]
}
