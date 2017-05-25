package mock

type HTTPAPIClient struct {
	GetFn   func(url string, resType interface{}) error
	GetCall int

	PostFn   func(url string, reqBody, resType interface{}) error
	PostCall int

	PutFn   func(url string, reqBody, resType interface{}) error
	PutCall int

	DeleteFn   func(url string, resType interface{}) error
	DeleteCall int

	PingFn   func() error
	PingCall int
}

func (c *HTTPAPIClient) Get(url string, resType interface{}) error {
	c.GetCall++

	return c.GetFn(url, resType)
}

func (c *HTTPAPIClient) Post(url string, reqBody, resType interface{}) error {
	c.PostCall++

	return c.PostFn(url, reqBody, resType)
}

func (c *HTTPAPIClient) Put(url string, reqBody, resType interface{}) error {
	c.PutCall++

	return c.PutFn(url, reqBody, resType)
}

func (c *HTTPAPIClient) Delete(url string, resType interface{}) error {
	c.DeleteCall++

	return c.DeleteFn(url, resType)
}

func (c *HTTPAPIClient) Ping() error {
	c.PingCall++

	return c.PingFn()
}
