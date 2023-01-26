package client

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestHistoryResource(t *testing.T) {
	c := New()
	assert.Equal(t, c.History(2021).Path, "/v1/hist_price_2021", "incorrect resource path")
	assert.Equal(t, c.History(2022).Path, "/v1/hist_price_2022", "incorrect resource path")
	assert.Equal(t, c.History(2023).Path, "/v1/hist_price_2023", "incorrect resource path")
}
