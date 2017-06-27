package ovh

import (
	go_ovh "github.com/ovh/go-ovh/ovh"

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

// Sync gathers all the existing redirection on your OVH domain.
func (h *DomainHandler) Sync() ([]*cochonou.Redirection, error) {
	domainRedirs, err := h.Client.GetAllRedirections()
	if err != nil {
		return nil, err
	}

	redirs := []*cochonou.Redirection{}

	for _, redir := range domainRedirs {
		redirs = append(redirs, redir.toRedirection())
	}

	return redirs, nil
}

func NewDomainHandler(domain string, ovhClient *go_ovh.Client) *DomainHandler {
	return &DomainHandler{
		Client: &Client{
			OVHWrapper: &HTTPAPIWrapper{
				Domain: domain,
				Client: ovhClient,
			},
		},
	}
}
