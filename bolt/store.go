package bolt

import "github.com/genesor/cochonou"

// ImageRedirectionStore is the implementation of cochonou.ImageRedirectionStore for BoltDB.
type ImageRedirectionStore struct {
}

// Save saves an ImageRedirection into the Bolt database.
func (s *ImageRedirectionStore) Save(redir *cochonou.ImageRedirection) error {
	return nil
}

// All fetches all the ImageRedirection from the Bolt database.
func (s *ImageRedirectionStore) All() ([]cochonou.ImageRedirection, error) {
	list := make([]cochonou.ImageRedirection, 0)

	return list, nil
}
