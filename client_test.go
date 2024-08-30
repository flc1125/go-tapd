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

const (
	apiUsername     = "tapd-username"
	apiPassword     = "tapd-password"
	successResponse = `{
  "status": 1,
  "data": {},
  "info": "success"
}`
)

var ctx = context.Background()

func createServerClient(t *testing.T, handler http.Handler) (*httptest.Server, *Client) {
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	client, err := NewClient(
		apiUsername, apiPassword,
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
		assert.Equal(t, apiUsername, username)
		assert.Equal(t, apiPassword, password)

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
  "status": 0,
  "data": {},
  "info": "error"
}`)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/error-response", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.Error(t, err)
	assert.True(t, IsErrorResponse(err))
}

func TestClient_NormalRequest(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/normal-request", r.URL.Path)

		// check header
		assert.Equal(t, defaultUserAgent, r.Header.Get("User-Agent"))

		// check basic auth
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, apiUsername, username)
		assert.Equal(t, apiPassword, password)

		fmt.Fprint(w, successResponse)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/normal-request", nil, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_WithRequestOption(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/request-option", r.URL.Path)

		// check header
		assert.Equal(t, "header-value", r.Header.Get("header-name"))
		assert.Equal(t, "headers-value", r.Header.Get("headers-name"))
		assert.Equal(t, "func-value", r.Header.Get("func-name"))
		assert.Contains(t, r.Header.Values("func-name-2"), "func-value-2")
		assert.Contains(t, r.Header.Values("func-name-2"), "func-value-3")
		assert.Equal(t, "test-user-agent", r.Header.Get("User-Agent"))

		// check basic auth
		username, password, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "test-username", username)
		assert.Equal(t, "test-password", password)

		fmt.Fprint(w, successResponse)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/request-option", nil, []RequestOption{
		WithRequestBasicAuth("test-username", "test-password"),
		WithRequestHeader("header-name", "header-value"),
		WithRequestHeaders(map[string]string{
			"headers-name": "headers-value",
		}),
		WithRequestHeaderFunc(func(header http.Header) {
			header.Set("func-name", "func-value")
			header.Add("func-name-2", "func-value-2")
			header.Add("func-name-2", "func-value-3")
		}),
		WithRequestUserAgent("test-user-agent"),
	})
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
