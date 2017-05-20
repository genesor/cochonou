package cochonou

import "errors"

var (
	// ErrSubDomainUsed is thrown when a a creation of an ImageRedirection fails
	// because the subdomain wished for is already used.
	ErrSubDomainUsed = errors.New("The subdomain is already used")
)

// ImageRedirection is the struct that represents a redirection for an image.
type ImageRedirection struct {
	ID        int
	SubDomain string
	URL       string
}

// ImageRedirectionStore is the interface used to store ImageRedirections
type ImageRedirectionStore interface {
	Save(redir *ImageRedirection) error
	All() ([]ImageRedirection, error)
}
