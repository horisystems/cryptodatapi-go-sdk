package client

import (
	"strconv"
	simplejson "github.com/bitly/go-simplejson"
)

type HistoryResource struct {
	Resource
}

func NewHistoryResource(year int) *HistoryResource {

	r := &HistoryResource{
		Resource:Resource{
			Name: "History",
			Path: "/v1/hist_price_" + strconv.Itoa(year),
		},
	}
	return r
}

func (r *HistoryResource) GetBySymbol(symbol string) (*simplejson.Json, error) {
	return getClient().http.request("GET", r.Path + "/symbol/" + symbol, nil)
}
