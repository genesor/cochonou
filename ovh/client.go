package ovh

// Client is the struct containing all the needed calls for the OVH API
type Client struct {
	OVHWrapper APIWrapper
}

// GetDomainRedirectionByName call the API to retrieve all the needed information for a subdomain.
func (c *Client) GetDomainRedirectionByName(name string) (*DomainRedirection, error) {
	id, err := c.OVHWrapper.GetDomainRedirectionID(name)
	if err != nil {
		return nil, err
	}

	subRedir, err := c.OVHWrapper.GetDomainRedirection(id)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}

// GetDomainRedirection call the API to retrieve all the needed information for a subdomain.
func (c *Client) GetDomainRedirection(id int) (*DomainRedirection, error) {
	subRedir, err := c.OVHWrapper.GetDomainRedirection(id)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}

// CreateDomainRedirection call the API to create a sub domain redirection..
func (c *Client) CreateDomainRedirection(subdomain string, target string) (*DomainRedirection, error) {

	redir := DomainRedirection{
		Type:      "visiblePermanent",
		SubDomain: subdomain,
		Target:    target,
	}

	ovhRedir, err := c.OVHWrapper.PostDomainRedirection(&redir)
	if err != nil {
		return nil, err
	}

	err = c.OVHWrapper.DomainRefreshDNSZone()

	return ovhRedir, err
}
