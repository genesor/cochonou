package mock

import "github.com/genesor/cochonou"

type DomainHandler struct {
	CreateDomainRedirectionFn   func(subDomain string, dest string) error
	CreateDomainRedirectionCall int

	SyncFn   func() ([]*cochonou.Redirection, error)
	SyncCall int
}

func (h *DomainHandler) CreateDomainRedirection(subDomain, dest string) error {
	h.CreateDomainRedirectionCall++

	return h.CreateDomainRedirectionFn(subDomain, dest)
}

func (h *DomainHandler) Sync() ([]*cochonou.Redirection, error) {
	h.SyncCall++

	return h.SyncFn()
}
