package ovh_test

import (
	"testing"

	go_ovh "github.com/ovh/go-ovh/ovh"

	"github.com/davecgh/go-spew/spew"
	"github.com/genesor/cochonou/ovh"
)

func TestGetSubDomain(t *testing.T) {
	apiClient, err := go_ovh.NewClient("ovh-eu", "D7U3sqg6x0UfoZRU", "90K4bROi6Y29AKIqFynXbjCJ5t3AK3YB", "BGirZ32wumxm5oJCTZieCBZgByq3YrQ8")
	if err != nil {
		t.Fatal(err)
	}

	ovhWrapper := &ovh.HTTPAPIWrapper{
		Client: apiClient,
		Domain: "sadoma.so",
	}

	client := &ovh.Client{
		OVHWrapper: ovhWrapper,
	}

	redir, err := client.GetSubDomainRedirectionByName("sefjsef")

	spew.Dump(redir, err)

	redir2, err := client.GetSubDomainRedirection(2309203912830912)

	spew.Dump(redir2, err)

	// t.Fail()

	// sub, err := client.GetSubDomain("lel")

	// require.Nil(t, oklm)
	// require.Nil(t, err)
}
