package cochonou_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/genesor/cochonou"
	"github.com/genesor/cochonou/mock"
)

func setupStored() (*cochonou.StoredDomainHandler, *mock.DomainHandler, *mock.RedirectionStore) {
	handler := &mock.DomainHandler{}
	store := &mock.RedirectionStore{}

	storedHandler := &cochonou.StoredDomainHandler{
		DomainHandler: handler,
		Store:         store,
	}

	return storedHandler, handler, store
}
func TestStoredDomainHandlerCreate(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		storedHandler, mockHandler, store := setupStored()

		mockHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "subtest", sub)
			require.Equal(t, "http://sadoma.so", target)

			return nil
		}
		store.SaveFn = func(redir *cochonou.Redirection) error {
			require.Equal(t, "subtest", redir.SubDomain)
			require.Equal(t, "http://sadoma.so", redir.URL)
			require.Equal(t, 0, redir.ID)

			return nil
		}

		err := storedHandler.CreateDomainRedirection("subtest", "http://sadoma.so")

		require.NoError(t, err)
		require.Equal(t, 1, mockHandler.CreateDomainRedirectionCall)
		require.Equal(t, 1, store.SaveCall)
	})

	t.Run("NOK - Sub Handler error", func(t *testing.T) {
		storedHandler, mockHandler, store := setupStored()

		mockHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "subtest", sub)
			require.Equal(t, "http://sadoma.so", target)

			return cochonou.ErrSubDomainAlreadyExists
		}

		err := storedHandler.CreateDomainRedirection("subtest", "http://sadoma.so")

		require.Error(t, err)
		require.Equal(t, cochonou.ErrSubDomainAlreadyExists, err)
		require.Equal(t, 1, mockHandler.CreateDomainRedirectionCall)
		require.Equal(t, 0, store.SaveCall)
	})

	t.Run("NOK - Store error exists", func(t *testing.T) {
		storedHandler, mockHandler, store := setupStored()

		mockHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "subtest", sub)
			require.Equal(t, "http://sadoma.so", target)

			return nil
		}
		store.SaveFn = func(redir *cochonou.Redirection) error {
			require.Equal(t, "subtest", redir.SubDomain)
			require.Equal(t, "http://sadoma.so", redir.URL)
			require.Equal(t, 0, redir.ID)

			return cochonou.ErrSubDomainUsed
		}

		err := storedHandler.CreateDomainRedirection("subtest", "http://sadoma.so")

		require.Error(t, err)
		require.Equal(t, cochonou.ErrSubDomainAlreadyExists, err)
		require.Equal(t, 1, mockHandler.CreateDomainRedirectionCall)
		require.Equal(t, 1, store.SaveCall)
	})

	t.Run("NOK - Store error", func(t *testing.T) {
		storedHandler, mockHandler, store := setupStored()

		mockHandler.CreateDomainRedirectionFn = func(sub, target string) error {
			require.Equal(t, "subtest", sub)
			require.Equal(t, "http://sadoma.so", target)

			return nil
		}
		store.SaveFn = func(redir *cochonou.Redirection) error {
			require.Equal(t, "subtest", redir.SubDomain)
			require.Equal(t, "http://sadoma.so", redir.URL)
			require.Equal(t, 0, redir.ID)

			return errors.New("some error")
		}

		err := storedHandler.CreateDomainRedirection("subtest", "http://sadoma.so")

		require.Error(t, err)
		require.NotEqual(t, cochonou.ErrSubDomainAlreadyExists, err)
		require.Equal(t, 1, mockHandler.CreateDomainRedirectionCall)
		require.Equal(t, 1, store.SaveCall)
	})
}
