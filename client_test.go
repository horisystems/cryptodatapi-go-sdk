package client

import (
	"os"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestApi(t *testing.T) {
	c := New()
	// test authentication
	token, ok := c.Auth(os.Getenv("USER"), os.Getenv("PASS"))
	t.Log("Token: " + token)
	assert.True(t, ok, "Authentication failed!")
	assert.NotEmpty(t, token, "Authentication failed!")
	// live price
	lp := c.LivePrice()
	// history 2021
	h21 := c.History(2021)
	endpoints := []*Resource{c.TopGainers(), c.TopLosers(), c.Derivatives(), c.Dex(), c.Lending(), c.Spot(), NewResource(lp.Name, lp.Path), NewResource(h21.Name, h21.Path) /*, c.News()*/}
	var res *simplejson.Json
	var err error
	var id string
	for _, endpoint := range endpoints {
		t.Log("begin to test " + endpoint.Name + "'s GetAll api")
		res, err = endpoint.GetAll()
		assert.Empty(t, err, endpoint.Name+" GetAll Api test failed")

		id = ""
		switch endpoint.Name {
		case "Dex":
			id = res.Get("exchange_dex").GetIndex(0).Get("id").MustString()
		case "Lending":
			id = res.Get("exchange_lending").GetIndex(0).Get("id").MustString()
		case "Spot":
			id = res.Get("exchange_spot").GetIndex(0).Get("id").MustString()
		case "Derivatives":
			id = res.Get("exchange_derivatives").GetIndex(0).Get("id").MustString()
		case "TopGainers":
			id = res.Get("crypto_gainers").GetIndex(0).Get("id").MustString()
		case "TopLosers":
			id = res.Get("crypto_losers").GetIndex(0).Get("id").MustString()
		case "News":
			id = res.Get("crypto_news").GetIndex(0).Get("id").MustString()
		case "History":
			id = res.Get("historical_price").GetIndex(0).Get("id").MustString()
		case "LivePrice":
			id = res.Get("crypto_price").GetIndex(0).Get("id").MustString()
		}
		if id == "" {
			continue
		}
		t.Log("begin to test " + endpoint.Name + "'s GetByID api, ID: " + id)
		_, err = endpoint.GetById(id)
		assert.Empty(t, err, endpoint.Name+" GetByID Api test failed")
	}

	t.Log("begin to test LivePrice's GetBySymbol api")
	_, err = lp.GetBySymbol("btc")
	assert.Empty(t, err, " LivePrice GetBySymbol  Api test failed")

	t.Log("begin to test History2021's GetBySymbol api")
	_, err = h21.GetBySymbol("btc")
	assert.Empty(t, err, " History2021 GetBySymbol  Api test failed")

}
