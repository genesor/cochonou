package mock

import (
	"github.com/genesor/cochonou/ovh"
)

type APIWrapper struct {
	GetDomainRedirectionIDFn   func(string) (int, error)
	GetDomainRedirectionIDCall int

	GetDomainRedirectionFn   func(int) (*ovh.DomainRedirection, error)
	GetDomainRedirectionCall int

	PostDomainRedirectionFn   func(*ovh.DomainRedirection) (*ovh.DomainRedirection, error)
	PostDomainRedirectionCall int

	DomainRefreshDNSZoneFn   func() error
	DomainRefreshDNSZoneCall int
}

func (w *APIWrapper) GetDomainRedirectionID(name string) (int, error) {
	w.GetDomainRedirectionIDCall++

	return w.GetDomainRedirectionIDFn(name)
}

func (w *APIWrapper) GetDomainRedirection(ID int) (*ovh.DomainRedirection, error) {
	w.GetDomainRedirectionCall++

	return w.GetDomainRedirectionFn(ID)
}

func (w *APIWrapper) PostDomainRedirection(subRedir *ovh.DomainRedirection) (*ovh.DomainRedirection, error) {
	w.PostDomainRedirectionCall++

	return w.PostDomainRedirectionFn(subRedir)
}

func (w *APIWrapper) DomainRefreshDNSZone() error {
	w.DomainRefreshDNSZoneCall++

	return w.DomainRefreshDNSZoneFn()
}
