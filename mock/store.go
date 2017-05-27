package mock

import "github.com/genesor/cochonou"

type RedirectionStore struct {
	SaveFn   func(redir *cochonou.Redirection) error
	SaveCall int
}

func (s *RedirectionStore) Save(redir *cochonou.Redirection) error {
	s.SaveCall++

	return s.SaveFn(redir)
}
