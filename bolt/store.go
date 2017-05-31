package bolt

import (
	"github.com/asdine/storm"

	"github.com/genesor/cochonou"
)

// RedirectionStore is the implementation of cochonou.RedirectionStore for BoltDB.
type RedirectionStore struct {
	DB storm.Node
}

// Save saves an Redirection into the Bolt database.
func (s *RedirectionStore) Save(redir *cochonou.Redirection) error {
	redirBolt := toBoltRedirection(redir)

	err := s.DB.Save(redirBolt)
	if err != nil {
		if err == storm.ErrAlreadyExists {
			return cochonou.ErrSubDomainUsed
		}

		return err
	}

	*redir = *fromBoltRedirection(redirBolt)

	return nil
}

// All fetches all the Redirection from the Bolt database.
// func (s *RedirectionStore) All() ([]cochonou.Redirection, error) {
// 	list := make([]cochonou.Redirection, 0)
//
// 	return list, nil
// }

// GetBySubDomain retrieves a redirection by its subdomain.
func (s *RedirectionStore) GetBySubDomain(subdomain string) (*cochonou.Redirection, error) {
	boltRedir := new(Redirection)

	err := s.DB.One("SubDomain", subdomain, boltRedir)
	if err != nil {
		if err == storm.ErrNotFound {
			return nil, cochonou.ErrNotFound
		}

		return nil, err
	}

	return fromBoltRedirection(boltRedir), nil
}
