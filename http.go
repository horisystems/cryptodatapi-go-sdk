package client

import (
	"net/url"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	BaseURL *url.URL
	resty   *resty.Client
	Token   string
}

func newHttpClient() *HttpClient {
	baseURL, _ := url.Parse("https://cryptodatapi.com")
	resty := resty.New()
	h := &HttpClient{
		BaseURL: baseURL,
		resty:   resty,
	}
	return h
}

func (h *HttpClient) request(method, path string, body interface{}) (*simplejson.Json, error) {

	rel := &url.URL{Path: path}
	u := h.BaseURL.ResolveReference(rel)

	req := h.resty.R()
	req.SetHeader("Accept", "application/json")
	if h.Token != "" {
		req.SetHeader("Authorization", "Bearer "+h.Token)
	}
	if body != nil {
		req.SetBody(body)
	}
	res, _ := req.Execute(method, u.String())
	return simplejson.NewJson(res.Body())
}
