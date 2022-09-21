package client

import (
	"testing"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestPriceResource(t *testing.T) {
	c := New()
	assert.Equal(t, c.LivePrice().Path, "/v1/crypto_price", "incorrect resource path")

}
