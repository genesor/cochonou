package bolt

import (
	"github.com/genesor/cochonou"
)

// ImageRedirection is the struct that represents a cochonou.ImageRedirection
// inside the BoltDB.
type ImageRedirection struct {
	ID        int    `storm:"id,increment"`
	SubDomain string `storm:"unique"`
	URL       string
}

func toBoltImageRedirection(img *cochonou.ImageRedirection) *ImageRedirection {
	return &ImageRedirection{
		ID:        img.ID,
		SubDomain: img.SubDomain,
		URL:       img.URL,
	}
}

func fromBoltImageRedirection(img *ImageRedirection) *cochonou.ImageRedirection {
	return &cochonou.ImageRedirection{
		ID:        img.ID,
		SubDomain: img.SubDomain,
		URL:       img.URL,
	}
}
