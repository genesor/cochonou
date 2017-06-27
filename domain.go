package cochonou

import (
	"errors"

	"github.com/ovh/go-ovh/ovh"
)

var (
	// ErrSubDomainAlreadyExists is used when a subdomain is already taken.
	ErrSubDomainAlreadyExists = errors.New("the subdomain cannot be created because it already exists")
)

// DomainHandler is the interface used to handle domain related operations.
type DomainHandler interface {
	CreateDomainRedirection(subDomain string, dest string) error
	Sync() ([]*Redirection, error)
}

// StoredDomainHandler is the implementation of DomainHandler
// With a layer that store the redirection in a store.
type StoredDomainHandler struct {
	DomainHandler DomainHandler
	OVHClient     ovh.Client
	Store         RedirectionStore
}

// CreateDomainRedirection calls another DomainHandler and save the redirection inside
// a store.
func (h *StoredDomainHandler) CreateDomainRedirection(subDomain string, dest string) error {
	_, err := h.Store.GetBySubDomain(subDomain)
	if err != ErrNotFound {
		if err == nil {
			return ErrSubDomainAlreadyExists
		}

		return err
	}

	err = h.DomainHandler.CreateDomainRedirection(subDomain, dest)
	if err != nil {
		return err
	}

	redir := Redirection{
		URL:       dest,
		SubDomain: subDomain,
	}

	err = h.Store.Save(&redir)
	if err != nil {
		if err == ErrSubDomainUsed {
			return ErrSubDomainAlreadyExists
		}

		return err
	}

	return nil
}

// Sync fetches all the domains from the provider and saves them inside the database
func (h *StoredDomainHandler) Sync() ([]*Redirection, error) {
	redirs, err := h.DomainHandler.Sync()
	if err != nil {
		return nil, err
	}

	for _, redir := range redirs {
		_, err := h.Store.GetBySubDomain(redir.SubDomain)
		if err == ErrNotFound {
			err := h.Store.Save(redir)
			if err != nil {
				return nil, err
			}
		}
	}

	return redirs, nil
}
