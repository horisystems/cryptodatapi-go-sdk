package client

type Client struct {
	http *HttpClient
}

var c *Client

func New() *Client {

	httpClient := newHttpClient()
	c = &Client{
		http: httpClient,
	}
	return c
}

func getClient() *Client {
	return c
}

func (c *Client) Auth(username, password string) (string, bool) {
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	json, err := c.http.request("POST", "/v1/login", data)
	if err != nil {
		return "", false
	}

	t, err := json.Get("Document").String()
	if t == "" || err != nil {
		return "", false
	}
	c.http.Token = t
	return t, true
}

func (c *Client) TopGainers() *Resource {
	r := NewResource("TopGainers", "/v1/gainers")
	return r
}

func (c *Client) TopLosers() *Resource {
	r := NewResource("TopLosers", "/v1/losers")
	return r
}

func (c *Client) Derivatives() *Resource {
	r := NewResource("Derivatives", "/v1/derivatives")
	return r
}

func (c *Client) Dex() *Resource {
	r := NewResource("Dex", "/v1/dex")
	return r
}

func (c *Client) Lending() *Resource {
	r := NewResource("Lending", "/v1/lending")
	return r
}

func (c *Client) Spot() *Resource {
	r := NewResource("Spot", "/v1/spot")
	return r
}

func (c *Client) News() *Resource {
	r := NewResource("News", "/v1/news")
	return r
}

func (c *Client) History(year int) *HistoryResource {
	r := NewHistoryResource(year)
	return r
}

func (c *Client) LivePrice() *PriceResource {
	r := NewPriceResource()
	return r
}
