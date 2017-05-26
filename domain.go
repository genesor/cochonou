package cochonou

import "errors"

var (
	// ErrSubDomainAlreadyExists is used when a subdomain is already taken.
	ErrSubDomainAlreadyExists = errors.New("the subdomain cannot be created because it already exists")
)

// DomainHandler is the interface used to handle domain related operations.
type DomainHandler interface {
	CreateDomainRedirection(subDomain string, dest string) error
	// GetAllRegisteredSubDomains()
}
