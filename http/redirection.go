package http

import (
	"github.com/genesor/cochonou"
)

// JSONRedirection is the struct that represents a cochonou.Redirection
// in JSON.
type JSONRedirection struct {
	ID        int    `json:"id"`
	SubDomain string `json:"sub_domain"`
	URL       string `json:"url"`
}

func toJSONRedirection(r *cochonou.Redirection) *JSONRedirection {
	return &JSONRedirection{
		ID:        r.ID,
		SubDomain: r.SubDomain,
		URL:       r.URL,
	}
}
