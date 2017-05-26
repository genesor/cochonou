package cochonou

// DomainHandler is the interface used to handle domain related operations.
type DomainHandler interface {
	CreateDomainRedirection(subDomain string, dest string) error
	// GetAllRegisteredSubDomains()
}
