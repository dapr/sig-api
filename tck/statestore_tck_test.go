package main

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"

	clients "github.com/darp/sig-api/tck/clients"
)

var statestoreName = "statestore"
var daprURLHTTP1 = "http://localhost:3500"

func Test_storeStateAPIs(t *testing.T) {

	ctx := context.Background()
	client, err := clients.NewClient(daprURLHTTP1)
	if err != nil {
		panic(err)
	}

	body := `[{ "key": "state_1", "value": "init"}]`

	resp, err := client.StoreStateWithBody(ctx, statestoreName, "application/json", bytes.NewReader([]byte(body)))

	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	resp, err = client.GetState(ctx, statestoreName, "state_1", &clients.GetStateParams{})
	if err != nil {
		panic(err)
	}

	var payload string
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "init", payload)

}
