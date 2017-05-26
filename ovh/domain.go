package ovh

// DomainHandler is the implementation of cochonou.DomainHandler for OVH.
type DomainHandler struct {
}

// CreateDomainRedirection creates a new subdomain redirection for your OVH domain
// using the API.
func (h *DomainHandler) CreateDomainRedirection(subDomain string, dest string) error {
	return nil
}
