package ovh

// Client is the struct containing all the needed calls for the OVH API
type Client struct {
	OVHWrapper APIWrapper
}

// GetSubDomainRedirectionByName call the API to retrieve all the needed information for a subdomain.
func (c *Client) GetSubDomainRedirectionByName(name string) (*SubDomainRedirection, error) {

	id, err := c.OVHWrapper.GetSubDomainRedirectionID(name)
	if err != nil {
		return nil, err
	}

	subRedir, err := c.OVHWrapper.GetSubDomainRedirection(id)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}

// GetSubDomainRedirection call the API to retrieve all the needed information for a subdomain.
func (c *Client) GetSubDomainRedirection(id int) (*SubDomainRedirection, error) {
	subRedir, err := c.OVHWrapper.GetSubDomainRedirection(id)
	if err != nil {
		return nil, err
	}

	return subRedir, nil
}
