package bolt

import (
	"github.com/genesor/cochonou"
)

// Redirection is the struct that represents a cochonou.Redirection
// inside the BoltDB.
type Redirection struct {
	ID        int    `storm:"id,increment"`
	SubDomain string `storm:"unique"`
	URL       string
}

func toBoltRedirection(img *cochonou.Redirection) *Redirection {
	return &Redirection{
		ID:        img.ID,
		SubDomain: img.SubDomain,
		URL:       img.URL,
	}
}

func fromBoltRedirection(img *Redirection) *cochonou.Redirection {
	return &cochonou.Redirection{
		ID:        img.ID,
		SubDomain: img.SubDomain,
		URL:       img.URL,
	}
}
