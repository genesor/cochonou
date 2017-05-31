package mock

import "github.com/genesor/cochonou"

type RedirectionStore struct {
	SaveFn   func(redir *cochonou.Redirection) error
	SaveCall int

	GetBySubDomainFn   func(string) (*cochonou.Redirection, error)
	GetBySubDomainCall int
}

func (s *RedirectionStore) Save(redir *cochonou.Redirection) error {
	s.SaveCall++

	return s.SaveFn(redir)
}

func (s *RedirectionStore) GetBySubDomain(subdomain string) (*cochonou.Redirection, error) {
	s.GetBySubDomainCall++

	return s.GetBySubDomainFn(subdomain)
}
