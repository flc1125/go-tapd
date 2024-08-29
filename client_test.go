package tapd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func createServerClient(t *testing.T, handler http.Handler) (*httptest.Server, *Client) {
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	client, err := NewClient(
		"tapd-username", "tapd-password",
		WithBaseURL(srv.URL),
		WithHTTPClient(NewRetryableHTTPClient(
			WithRetryableHTTPClientLogger(log.New(os.Stderr, "", log.LstdFlags)),
		)),
	)
	assert.NoError(t, err)

	return srv, client
}

func loadData(t *testing.T, filepath string) []byte {
	content, err := os.ReadFile(filepath)
	assert.NoError(t, err)
	return content
}

func TestClient_BasicAuth(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/basic-auth", r.URL.Path)

		// check basic auth
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "tapd-username", username)
		assert.Equal(t, "tapd-password", password)

		fmt.Fprint(w, `{
  "status": 1,
  "data": {},
  "info": "success"
}`)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/basic-auth", nil, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_ErrorResponse(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/error-response", r.URL.Path)

		fmt.Fprint(w, `{
  "status": 429,
  "data": {},
  "info": "To many requests, api account brookechen max request rates is 6000req/10min"
}`)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/error-response", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.Error(t, err)
	assert.True(t, IsErrorResponse(err))
}
