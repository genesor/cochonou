package ovh

import (
	"errors"
	"path"
	"strconv"

	"github.com/genesor/cochonou"
)

var (
	// ErrNoResult is used when an API request ends up with no result.
	ErrNoResult = errors.New("no result found")
	// ErrNonUniqueResult is used when a single result was expected and more
	// were given.
	ErrNonUniqueResult = errors.New("non unique result")
)

// DomainRedirection represents an OVH subdomain redirection.
type DomainRedirection struct {
	ID          int    `json:"id,omitempty"`
	Zone        string `json:"zone,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Target      string `json:"target"`
	SubDomain   string `json:"subDomain"`
	Type        string `json:"type"` // invisible, visible, visiblePermanent
	Title       string `json:"title,omitempty"`
}

func (dr *DomainRedirection) toRedirection() *cochonou.Redirection {
	return &cochonou.Redirection{
		SubDomain: dr.SubDomain,
		URL:       dr.Target,
	}
}

// APIWrapper is the interface of a bridge between OVH and differents structs.
type APIWrapper interface {
	GetDomainRedirectionIDs() ([]int, error)
	GetDomainRedirectionID(name string) (int, error)
	GetDomainRedirection(ID int) (*DomainRedirection, error)
	PostDomainRedirection(*DomainRedirection) (*DomainRedirection, error)
	DomainRefreshDNSZone() error
}

// HTTPAPIWrapper is the HTTP wrapper for the OVH API
type HTTPAPIWrapper struct {
	Client HTTPAPIClient
	Domain string
}

// GetDomainRedirectionID fetches the ID of a redirection from its subdomain value.
func (w *HTTPAPIWrapper) GetDomainRedirectionID(name string) (int, error) {
	redirectionIDs := make([]int, 0)

	err := w.Client.Get(path.Join("/domain/zone/", w.Domain, "/redirection?subDomain="+name), &redirectionIDs)
	if err != nil {
		return 0, err
	}

	if len(redirectionIDs) == 0 {
		return 0, ErrNoResult
	} else if len(redirectionIDs) != 1 {
		return 0, ErrNonUniqueResult
	}

	return redirectionIDs[0], nil
}

// GetDomainRedirectionIDs fetches the ID of a redirection from its subdomain value.
func (w *HTTPAPIWrapper) GetDomainRedirectionIDs() ([]int, error) {
	redirectionIDs := make([]int, 0)

	err := w.Client.Get(path.Join("/domain/zone/", w.Domain, "/redirection"), &redirectionIDs)
	if err != nil {
		return nil, err
	}

	return redirectionIDs, nil
}

// GetDomainRedirection fetches all the data for a DomainRedirection from its ID.
func (w *HTTPAPIWrapper) GetDomainRedirection(ID int) (*DomainRedirection, error) {
	subRedir := new(DomainRedirection)

	err := w.Client.Get(path.Join("/domain/zone/", w.Domain, "/redirection/", strconv.Itoa(ID)), subRedir)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}

// PostDomainRedirection fetches all the data for a DomainRedirection from its ID.
func (w *HTTPAPIWrapper) PostDomainRedirection(subRedir *DomainRedirection) (*DomainRedirection, error) {
	ovhRedir := new(DomainRedirection)

	err := w.Client.Post(path.Join("/domain/zone/", w.Domain, "/redirection"), subRedir, ovhRedir)
	if err != nil {
		return nil, err
	}

	return ovhRedir, nil
}

// DomainRefreshDNSZone ask to perform a refresh of the DNS Zone for the domain.
func (w *HTTPAPIWrapper) DomainRefreshDNSZone() error {
	err := w.Client.Post(path.Join("/domain/zone/", w.Domain, "/refresh"), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// HTTPAPIClient is the interface for an HTTP Client using an API
type HTTPAPIClient interface {
	Get(url string, resType interface{}) error
	Post(url string, reqBody, resType interface{}) error
	Put(url string, reqBody, resType interface{}) error
	Delete(url string, resType interface{}) error
	Ping() error
}
