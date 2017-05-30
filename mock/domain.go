package mock

type DomainHandler struct {
	CreateDomainRedirectionFn   func(subDomain string, dest string) error
	CreateDomainRedirectionCall int
}

func (h *DomainHandler) CreateDomainRedirection(subDomain, dest string) error {
	h.CreateDomainRedirectionCall++

	return h.CreateDomainRedirectionFn(subDomain, dest)
}
