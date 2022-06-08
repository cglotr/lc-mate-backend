package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/hooligram/kifu"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Setenv("PORT", "9000")
	main()
	url := "http://localhost:9000/ping"
	resp, err := http.Get(url)
	for err != nil {
		resp, err = http.Get(url)
		kifu.Warn("Retrying since webserver might not be online yet...")
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, `{"message":"pong"}`, string(bytes))
}
