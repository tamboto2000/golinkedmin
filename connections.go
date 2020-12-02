package golinkedmin

import "github.com/tamboto2000/golinkedin"

// Connections store profile connections.
// It is highly recommended to not modify Start and Count value
type Connections struct {
	Profiles []Profile `json:"profiles,omitempty"`
	Start    int       `json:"start,omitempty"`
	Count    int       `json:"count,omitempty"`
	Total    int       `json:"total,omitempty"`

	pnode      *golinkedin.PeopleNode
	ln         *Linkedin
	stopCursor bool
	err        error
}

// Next fetch next page of connections
func (cn *Connections) Next() bool {
	if cn.stopCursor {
		return false
	}

	cn.pnode.Paging.Start = cn.Start
	cn.pnode.Paging.Count = cn.Count

	miniProfs := make([]Profile, 0)
	for cn.pnode.Next() {
		for _, elm := range cn.pnode.Elements {
			if elm.Type == golinkedin.TypeSearchHits {
				for _, innerElm := range elm.Elements {
					miniProfs = append(miniProfs, *composeMiniProfile(innerElm.Image.Attributes[0].MiniProfile))
				}

				break
			}
		}
	}

	if err := cn.pnode.Error(); err != nil {
		cn.err = err
		return false
	}

	if len(miniProfs) < cn.pnode.Paging.Count {
		cn.stopCursor = true
	}

	cn.Profiles = miniProfs
	cn.Start = cn.pnode.Paging.Start
	cn.Count = cn.pnode.Paging.Count
	cn.Total = cn.pnode.Paging.Total

	return true
}

// SetLinkedin set new client
func (cn *Connections) SetLinkedin(ln *Linkedin) {
	cn.pnode.SetLinkedin(ln.Linkedin)
	cn.ln = ln
}

// Error return error
func (cn *Connections) Error() error {
	if cn.err != nil {
		return parseErrMsg(cn.err.Error())
	}

	return nil
}
