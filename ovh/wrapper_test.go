package ovh_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Genesor/cochonou/ovh"
	"github.com/genesor/cochonou/ovh/mock"
)

func setup() (*ovh.HTTPAPIWrapper, *mock.HTTPAPIClient) {
	client := &mock.HTTPAPIClient{}
	wrapper := &ovh.HTTPAPIWrapper{
		Domain: "test.com",
		Client: client,
	}

	return wrapper, client
}

func TestWrapperGetSubDomainRedirectionID(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirectionID")
			}
			*returnValue = []int{1234}

			return nil
		}

		ID, err := wrapper.GetSubDomainRedirectionID("subtest")

		require.NoError(t, err)
		require.Equal(t, 1234, ID)
	})

	t.Run("NOK - No result", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirectionID")
			}
			*returnValue = []int{}

			return nil
		}

		ID, err := wrapper.GetSubDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, ovh.ErrNoResult, err)
		require.Equal(t, 0, ID)
	})

	t.Run("NOK - Multi result", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirectionID")
			}
			*returnValue = []int{1234, 5678}

			return nil
		}

		ID, err := wrapper.GetSubDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, ovh.ErrNonUniqueResult, err)
		require.Equal(t, 0, ID)
	})

	t.Run("NOK - Error API", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			_, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirectionID")
			}

			return errors.New("some error")
		}

		ID, err := wrapper.GetSubDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, 0, ID)
	})
}

func TestWrapperGetSubDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection/1234", url)

			returnValue, ok := resType.(*ovh.SubDomainRedirection)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirection")
			}

			*returnValue = ovh.SubDomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}

			return nil
		}

		subRedir, err := wrapper.GetSubDomainRedirection(1234)

		require.NoError(t, err)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
	})

	t.Run("NOK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setup()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection/1234", url)

			_, ok := resType.(*ovh.SubDomainRedirection)
			if !ok {
				t.Error("Wrong type given by GetSubDomainRedirection")
			}

			return errors.New("some error")
		}

		subRedir, err := wrapper.GetSubDomainRedirection(1234)

		require.Error(t, err)
		require.Nil(t, subRedir)
	})
}
