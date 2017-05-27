package ovh

import (
	"github.com/genesor/cochonou"
)

// DomainHandler is the implementation of cochonou.DomainHandler for OVH.
type DomainHandler struct {
	Client *Client
}

// CreateDomainRedirection creates a new subdomain redirection for your OVH domain
// using the API.
func (h *DomainHandler) CreateDomainRedirection(subDomain string, dest string) error {
	_, err := h.Client.GetDomainRedirectionByName(subDomain)
	if err != ErrNoResult {
		if err != nil {
			return err
		}

		return cochonou.ErrSubDomainAlreadyExists
	}

	_, err = h.Client.CreateDomainRedirection(subDomain, dest)

	return err
}
