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

func TestGetSubDomainRedirectionByName(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetSubDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 1234, nil
		}

		wrapper.GetSubDomainRedirectionFn = func(ID int) (*ovh.SubDomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return &ovh.SubDomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}, nil
		}

		subRedir, err := client.GetSubDomainRedirectionByName("subtest")

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionIDCall)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionCall)
	})

	t.Run("NOK - Doesn't exist", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetSubDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 0, ovh.ErrNoResult
		}

		wrapper.GetSubDomainRedirectionFn = func(ID int) (*ovh.SubDomainRedirection, error) {

			return nil, errors.New("unexpected call")
		}

		subRedir, err := client.GetSubDomainRedirectionByName("subtest")

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, ovh.ErrNoResult, err)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionIDCall)
		require.Equal(t, 0, wrapper.GetSubDomainRedirectionCall)
	})

	t.Run("NOK - Error", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetSubDomainRedirectionIDFn = func(name string) (int, error) {
			require.Equal(t, "subtest", name)

			return 1234, nil
		}

		wrapper.GetSubDomainRedirectionFn = func(ID int) (*ovh.SubDomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return nil, errors.New("some error")
		}

		subRedir, err := client.GetSubDomainRedirectionByName("subtest")

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionIDCall)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionCall)
	})
}

func TestGetSubDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetSubDomainRedirectionFn = func(ID int) (*ovh.SubDomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return &ovh.SubDomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}, nil
		}

		subRedir, err := client.GetSubDomainRedirection(1234)

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionCall)
	})

	t.Run("NOK - Error", func(t *testing.T) {
		client, wrapper := setupClient()

		wrapper.GetSubDomainRedirectionFn = func(ID int) (*ovh.SubDomainRedirection, error) {
			require.Equal(t, 1234, ID)

			return nil, errors.New("some error")
		}

		subRedir, err := client.GetSubDomainRedirection(1234)

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, wrapper.GetSubDomainRedirectionCall)
	})
}
