package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"

	clients "github.com/salaboy/sig-api/tck/clients"
)

var statestoreName = "statestore"
var daprURLHTTP1 = "http://localhost:3500"

func Test_storeStateAPIsBasic(t *testing.T) {

	ctx := context.Background()
	client, err := clients.NewClient(daprURLHTTP1)
	if err != nil {
		panic(err)
	}

	body := `[{ "key": "state_1", "value": "init"}]`

	resp, err := client.PostStoreStateWithBody(ctx, statestoreName, "application/json", bytes.NewReader([]byte(body)))

	if err != nil {
		assert.Fail(t, "", err)
	}

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	resp, err = client.GetState(ctx, statestoreName, "state_1", &clients.GetStateParams{})
	if err != nil {
		assert.Fail(t, "", err)
	}

	var payload string
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		assert.Fail(t, "", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "init", payload)

	resp, err = client.DeleteState(ctx, statestoreName, "state_1")
	if err != nil {
		assert.Fail(t, "", err)
	}
	assert.Equal(t, 204, resp.StatusCode)

	resp, err = client.GetState(ctx, statestoreName, "state_1", &clients.GetStateParams{})
	if err != nil {
		assert.Fail(t, "", err)
	}
	var emptyPayload string
	err = json.NewDecoder(resp.Body).Decode(emptyPayload)

	assert.Equal(t, io.EOF, err)
}

func Test_bulkStoreStateAPIs(t *testing.T) {

}
