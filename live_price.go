package client

import (
	simplejson "github.com/bitly/go-simplejson"
)

type PriceResource struct {
	Resource
}

func NewPriceResource() *PriceResource {

	r := &PriceResource{
		Resource: Resource{
			Name: "LivePrice",
			Path: "/v1/crypto_price",
		},
	}
	return r
}


func (r *PriceResource) GetBySymbol(symbol string) (*simplejson.Json, error) {
	return getClient().http.request("GET", r.Path + "/symbol/" + symbol, nil)
}
