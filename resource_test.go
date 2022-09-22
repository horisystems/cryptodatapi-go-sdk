package client

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestResourcePath(t *testing.T) {
	c := New()

	assert.Equal(t, c.TopGainers().Path, "/v1/gainers", "incorrect resource path")
	assert.Equal(t, c.TopLosers().Path, "/v1/losers", "incorrect resource path")
	assert.Equal(t, c.Derivatives().Path, "/v1/derivatives", "incorrect resource path")
	assert.Equal(t, c.Dex().Path, "/v1/dex", "incorrect resource path")
	assert.Equal(t, c.Lending().Path, "/v1/lending", "incorrect resource path")
	assert.Equal(t, c.Spot().Path, "/v1/spot", "incorrect resource path")
	assert.Equal(t, c.News().Path, "/v1/news", "incorrect resource path")

}
