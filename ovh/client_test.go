package ovh_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/genesor/cochonou/ovh"
	"github.com/genesor/cochonou/ovh/mock"
)

func setupClient() (*ovh.Client, *mock.APIWrapper) {
	ovhWrapper := &mock.APIWrapper{}

	client := &ovh.Client{
		OVHWrapper: ovhWrapper,
	}

	return client, ovhWrapper
}

func TestGetDomainRedirectionByName(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 1234, nil
		}

		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return &ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}, nil
		}

		subRedir, err := client.GetDomainRedirectionByName("subtest")

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 1, wrapper.GetDomainRedirectionCall)
	})

	t.Run("NOK - Doesn't exist", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 0, ovh.ErrNoResult
		}

		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {

			return nil, errors.New("unexpected call")
		}

		subRedir, err := client.GetDomainRedirectionByName("subtest")

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, ovh.ErrNoResult, err)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 0, wrapper.GetDomainRedirectionCall)
	})

	t.Run("NOK - Error", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 1234, nil
		}

		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return nil, errors.New("some error")
		}

		subRedir, err := client.GetDomainRedirectionByName("subtest")

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, wrapper.GetDomainRedirectionIDCall)
		require.Equal(t, 1, wrapper.GetDomainRedirectionCall)
	})
}

func TestGetDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return &ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}, nil
		}

		subRedir, err := client.GetDomainRedirection(1234)

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, wrapper.GetDomainRedirectionCall)
	})

	t.Run("NOK - Error", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetDomainRedirectionFn = func(ID int) (*ovh.DomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return nil, errors.New("some error")
		}

		subRedir, err := client.GetDomainRedirection(1234)

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, wrapper.GetDomainRedirectionCall)
	})
}

func TestCreateDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		client, wrapper := setupClient()

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

		redir, err := client.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.NoError(t, err)
		require.NotNil(t, redir)
		require.Equal(t, 1234, redir.ID)
		require.Equal(t, "test.com", redir.Zone)
		require.Equal(t, "http://sadoma.so", redir.Target)
		require.Equal(t, "testcreate", redir.SubDomain)
		require.Equal(t, "visiblePermanent", redir.Type)
		require.Equal(t, 1, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 1, wrapper.DomainRefreshDNSZoneCall)
	})

	t.Run("NOK - Error create", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.PostDomainRedirectionFn = func(redir *ovh.DomainRedirection) (*ovh.DomainRedirection, error) {
			require.Equal(t, "http://sadoma.so", redir.Target)
			require.Equal(t, "visiblePermanent", redir.Type)
			require.Equal(t, "testcreate", redir.SubDomain)

			return nil, errors.New("some error")
		}
		wrapper.DomainRefreshDNSZoneFn = func() error {
			return errors.New("unexpected call")
		}

		redir, err := client.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.Error(t, err)
		require.Nil(t, redir)
		require.Equal(t, 1, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 0, wrapper.DomainRefreshDNSZoneCall)
	})

	t.Run("NOK - Error refresh", func(t *testing.T) {
		client, wrapper := setupClient()

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
			return errors.New("some error")
		}

		redir, err := client.CreateDomainRedirection("testcreate", "http://sadoma.so")

		require.Error(t, err)
		require.NotNil(t, redir)
		require.Equal(t, 1234, redir.ID)
		require.Equal(t, "test.com", redir.Zone)
		require.Equal(t, "http://sadoma.so", redir.Target)
		require.Equal(t, "testcreate", redir.SubDomain)
		require.Equal(t, "visiblePermanent", redir.Type)
		require.Equal(t, 1, wrapper.PostDomainRedirectionCall)
		require.Equal(t, 1, wrapper.DomainRefreshDNSZoneCall)
	})
}
