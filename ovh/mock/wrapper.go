package mock

import (
	"github.com/genesor/cochonou/ovh"
)

type APIWrapper struct {
	GetSubDomainRedirectionIDFn   func(string) (int, error)
	GetSubDomainRedirectionIDCall int

	GetSubDomainRedirectionFn   func(int) (*ovh.SubDomainRedirection, error)
	GetSubDomainRedirectionCall int
}

func (w *APIWrapper) GetSubDomainRedirectionID(name string) (int, error) {
	w.GetSubDomainRedirectionIDCall++

	return w.GetSubDomainRedirectionIDFn(name)
}

func (w *APIWrapper) GetSubDomainRedirection(ID int) (*ovh.SubDomainRedirection, error) {
	w.GetSubDomainRedirectionCall++

	return w.GetSubDomainRedirectionFn(ID)
}
