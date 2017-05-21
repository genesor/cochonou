package ovh

// DomainHandler is the implementation of cochonou.DomainHandler for OVH.
type DomainHandler struct {
}

// CreateSubDomainRedirection creates a new subdomain redirection for your OVH domain
// using the API.
func (h *DomainHandler) CreateSubDomainRedirection(subDomain string, dest string) error {
	return nil
}
