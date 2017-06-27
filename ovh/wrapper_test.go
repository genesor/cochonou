package ovh_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/genesor/cochonou/ovh"
	"github.com/genesor/cochonou/ovh/mock"
)

func setupWrapper() (*ovh.HTTPAPIWrapper, *mock.HTTPAPIClient) {
	client := &mock.HTTPAPIClient{}
	wrapper := &ovh.HTTPAPIWrapper{
		Domain: "test.com",
		Client: client,
	}

	return wrapper, client
}

func TestWrapperGetDomainRedirectionID(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}
			*returnValue = []int{1234}

			return nil
		}

		ID, err := wrapper.GetDomainRedirectionID("subtest")

		require.NoError(t, err)
		require.Equal(t, 1234, ID)
		require.Equal(t, 1, client.GetCall)
	})

	t.Run("NOK - No result", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}
			*returnValue = []int{}

			return nil
		}

		ID, err := wrapper.GetDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, ovh.ErrNoResult, err)
		require.Equal(t, 0, ID)
		require.Equal(t, 1, client.GetCall)
	})

	t.Run("NOK - Multi result", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}
			*returnValue = []int{1234, 5678}

			return nil
		}

		ID, err := wrapper.GetDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, ovh.ErrNonUniqueResult, err)
		require.Equal(t, 0, ID)
		require.Equal(t, 1, client.GetCall)
	})

	t.Run("NOK - Error API", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection?subDomain=subtest", url)

			_, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}

			return errors.New("some error")
		}

		ID, err := wrapper.GetDomainRedirectionID("subtest")

		require.Error(t, err)
		require.Equal(t, 0, ID)
		require.Equal(t, 1, client.GetCall)
	})
}

func TestWrapperGetDomainRedirectionIDs(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection", url)

			returnValue, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}
			*returnValue = []int{1234}

			return nil
		}

		IDs, err := wrapper.GetDomainRedirectionIDs()

		require.NoError(t, err)
		require.Equal(t, 1, len(IDs))
		require.Equal(t, 1234, IDs[0])
		require.Equal(t, 1, client.GetCall)
	})

	t.Run("NOK - Error API", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection", url)

			_, ok := resType.(*[]int)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirectionID")
			}

			return errors.New("some error")
		}

		_, err := wrapper.GetDomainRedirectionIDs()

		require.Error(t, err)
		require.Equal(t, 1, client.GetCall)
	})
}

func TestWrapperGetDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection/1234", url)

			returnValue, ok := resType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirection")
			}

			*returnValue = ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}

			return nil
		}

		subRedir, err := wrapper.GetDomainRedirection(1234)

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, client.GetCall)
	})

	t.Run("NOK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.GetFn = func(url string, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection/1234", url)

			_, ok := resType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by GetDomainRedirection")
			}

			return errors.New("some error")
		}

		subRedir, err := wrapper.GetDomainRedirection(1234)

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, client.GetCall)
	})
}

func TestWrapperPostDomainRedirection(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.PostFn = func(url string, reqType interface{}, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection", url)

			reqRedir, ok := reqType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by PostDomainRedirection")
			}

			require.Equal(t, 0, reqRedir.ID)
			require.Equal(t, "subtest", reqRedir.SubDomain)
			require.Equal(t, "visiblePermanent", reqRedir.Type)

			returnValue, ok := resType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by PostDomainRedirection")
			}

			*returnValue = ovh.DomainRedirection{
				ID:        1234,
				Zone:      "test.com",
				Target:    "http://sadoma.so",
				SubDomain: "subtest",
				Type:      "visiblePermanent",
			}

			return nil
		}

		redir := ovh.DomainRedirection{
			Zone:      "test.com",
			Target:    "http://sadoma.so",
			SubDomain: "subtest",
			Type:      "visiblePermanent",
		}

		subRedir, err := wrapper.PostDomainRedirection(&redir)

		require.NoError(t, err)
		require.NotNil(t, subRedir)
		require.Equal(t, 1234, subRedir.ID)
		require.Equal(t, "test.com", subRedir.Zone)
		require.Equal(t, "http://sadoma.so", subRedir.Target)
		require.Equal(t, "subtest", subRedir.SubDomain)
		require.Equal(t, "visiblePermanent", subRedir.Type)
		require.Equal(t, 1, client.PostCall)
	})

	t.Run("NOK", func(t *testing.T) {
		t.Parallel()

		wrapper, client := setupWrapper()

		client.PostFn = func(url string, reqType interface{}, resType interface{}) error {
			require.Equal(t, "/domain/zone/test.com/redirection", url)

			reqRedir, ok := reqType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by PostDomainRedirection")
			}

			require.Equal(t, 0, reqRedir.ID)
			require.Equal(t, "subtest", reqRedir.SubDomain)
			require.Equal(t, "visiblePermanent", reqRedir.Type)

			_, ok = resType.(*ovh.DomainRedirection)
			if !ok {
				t.Error("Wrong type given by PostDomainRedirection")
			}

			return errors.New("some error")
		}

		redir := ovh.DomainRedirection{
			Zone:      "test.com",
			Target:    "http://sadoma.so",
			SubDomain: "subtest",
			Type:      "visiblePermanent",
		}

		subRedir, err := wrapper.PostDomainRedirection(&redir)

		require.Error(t, err)
		require.Nil(t, subRedir)
		require.Equal(t, 1, client.PostCall)
	})
}
