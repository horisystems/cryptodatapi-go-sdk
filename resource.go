package client

import (
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type Resource struct {
	Name string
	Path string
}

func NewResource(name, path string) *Resource {

	r := &Resource{
		Name: name,
		Path: path,
	}
	return r
}

func (r *Resource) GetAll() (*simplejson.Json, error) {
	return getClient().http.request("GET", r.Path, nil)
}

func (r *Resource) GetById(id string) (*simplejson.Json, error) {
	return getClient().http.request("GET", formatPath(r.Path, id), nil)
}

func (r *Resource) Create(data map[string]interface{}) {
	getClient().http.request("POST", r.Path, data)
}

func (r *Resource) DeleteById(id string) {
	getClient().http.request("DELETE", formatPath(r.Path, id), nil)
}

func (r *Resource) UpdateById(id string, data map[string]interface{}) {
	getClient().http.request("PUT", formatPath(r.Path, id), data)
}

// func (r *PriceResource) GetBySymbol(symbol string) (*simplejson.Json, error) {
// 	return getClient().http.request("GET", r.Path + "/symbol/" + symbol, nil)
// }

func formatPath(paths ...string) string {
	var str string
	for _, path := range paths {
		str += strings.Trim(path, "/") + "/"
	}
	return strings.TrimRight(str, "/")
}
