package cochonou

import "errors"

var (
	// ErrSubDomainUsed is thrown when a a creation of an Redirection fails
	// because the subdomain wished for is already used.
	ErrSubDomainUsed = errors.New("The subdomain is already used")
	// ErrNotFound is thrown when no redirection is found.
	ErrNotFound = errors.New("no redirection found")
)

// Redirection is the struct that represents a redirection.
type Redirection struct {
	ID        int
	SubDomain string
	URL       string
}

// RedirectionStore is the interface used to store Redirections
type RedirectionStore interface {
	Save(redir *Redirection) error
	GetBySubDomain(string) (*Redirection, error)
	GetAll() ([]*Redirection, error)
}
