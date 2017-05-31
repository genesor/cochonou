package cochonou

import "errors"

var (
	// ErrSubDomainAlreadyExists is used when a subdomain is already taken.
	ErrSubDomainAlreadyExists = errors.New("the subdomain cannot be created because it already exists")
)

// DomainHandler is the interface used to handle domain related operations.
type DomainHandler interface {
	CreateDomainRedirection(subDomain string, dest string) error
}

// StoredDomainHandler is the implementation of DomainHandler
// With a layer that store the redirection in a store.
type StoredDomainHandler struct {
	DomainHandler DomainHandler
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
