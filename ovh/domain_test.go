package ovh_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Genesor/cochonou"
	"github.com/genesor/cochonou/ovh"
	"github.com/genesor/cochonou/ovh/mock"
)

func setupDomainHandler() (*ovh.DomainHandler, *mock.APIWrapper) {
	ovhWrapper := &mock.APIWrapper{}

	domainHandler := &ovh.DomainHandler{
		Client: &ovh.Client{
			OVHWrapper: ovhWrapper,
		},
	}

	return domainHandler, ovhWrapper
}

func TestHandlerCreateDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		handler, wrapper := setupDomainHandler()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "testcreate", name)

			return 0, ovh.ErrNoResult
		}
		wrapper.PostDomainRedirectionFn = func(redir *ovh.DomainRedirection) (*ovh.DomainRedirection, error) {
			require.Equal(t, "http://sadoma.so", redir.Target)
			require.Equal(t, "visiblePermanent", redir.Type)
			require.Equal(t, "testcreate", redir.SubDomain)

			return &ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "testcreate",
				Type:      "visiblePermanent",
			}, nil
		}
		wrapper.DomainRefreshDNSZoneFn = func() error {
			return nil
		}

		err := handler.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.NoError(t, err)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 0, wrapper.GetDomainRedirectionCall)
		require.Equal(t, 1, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 1, wrapper.DomainRefreshDNSZoneCall)
	})

	t.Run("NOK - Domain exists", func(t *testing.T) {
		handler, wrapper := setupDomainHandler()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "testcreate", name)

			return 1234, nil
		}
		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return &ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "testcreate",
				Type:      "visiblePermanent",
			}, nil
		}

		err := handler.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.Error(t, err)
		require.Equal(t, cochonou.ErrSubDomainAlreadyExists, err)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 1, wrapper.GetDomainRedirectionCall)
		require.Equal(t, 0, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 0, wrapper.DomainRefreshDNSZoneCall)
	})

	t.Run("NOK - Error create", func(t *testing.T) {
		handler, wrapper := setupDomainHandler()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "testcreate", name)

			return 0, ovh.ErrNoResult
		}
		wrapper.PostDomainRedirectionFn = func(redir *ovh.DomainRedirection) (*ovh.DomainRedirection, error) {
			require.Equal(t, "http://sadoma.so", redir.Target)
			require.Equal(t, "visiblePermanent", redir.Type)
			require.Equal(t, "testcreate", redir.SubDomain)

			return nil, errors.New("some error")
		}

		err := handler.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.Error(t, err)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 0, wrapper.GetDomainRedirectionCall)
		require.Equal(t, 1, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 0, wrapper.DomainRefreshDNSZoneCall)
	})
}
