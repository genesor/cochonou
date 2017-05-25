package ovh

import (
	"errors"
	"path"
	"strconv"
)

var (
	// ErrNoResult is used when an API request ends up with no result.
	ErrNoResult = errors.New("no result found")
	// ErrNonUniqueResult is used when a single result was expected and more
	// were given.
	ErrNonUniqueResult = errors.New("non unique result")
)

// APIWrapper is the interface of a bridge between OVH and differents structs.
type APIWrapper interface {
	GetSubDomainRedirectionID(name string) (int, error)
	GetSubDomainRedirection(ID int) (*SubDomainRedirection, error)
}

// HTTPAPIWrapper is the HTTP wrapper for the OVH API
type HTTPAPIWrapper struct {
	Client HTTPAPIClient
	Domain string
}

// GetSubDomainRedirectionID fetches this ID of a redirection from its subdomain value.
func (w *HTTPAPIWrapper) GetSubDomainRedirectionID(name string) (int, error) {
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

// SubDomainRedirection represents an OVH subdomain redirection.
type SubDomainRedirection struct {
	ID          int    `json:"id"`
	Zone        string `json:"zone"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Target      string `json:"target"`
	SubDomain   string `json:"subDomain"`
	Type        string `json:"type"` // invisible, visible, visiblePermanent
	Title       string `json:"title"`
}

func (w *HTTPAPIWrapper) GetSubDomainRedirection(ID int) (*SubDomainRedirection, error) {
	subRedir := new(SubDomainRedirection)

	err := w.Client.Get(path.Join("/domain/zone/", w.Domain, "/redirection/", strconv.Itoa(ID)), subRedir)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}

// HTTPAPIClient is the interface for an HTTP Client using an API
type HTTPAPIClient interface {
	Get(url string, resType interface{}) error
	Post(url string, reqBody, resType interface{}) error
	Put(url string, reqBody, resType interface{}) error
	Delete(url string, resType interface{}) error
	Ping() error
}
