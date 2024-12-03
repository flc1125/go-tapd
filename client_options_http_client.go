package tapd

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type RetryableHTTPClientOption func(client *retryablehttp.Client)

func WithRetryableHTTPClientLogger(logger retryablehttp.Logger) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.Logger = logger
	}
}

func WithRetryableHTTPClientRetryWaitMin(min time.Duration) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.RetryWaitMin = min
	}
}

func WithRetryableHTTPClientRetryWaitMax(max time.Duration) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.RetryWaitMax = max
	}
}

func WithRetryableHTTPClientRetryMax(max int) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.RetryMax = max
	}
}

func WithRetryableHTTPClientCheckRetry(checkRetry retryablehttp.CheckRetry) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.CheckRetry = checkRetry
	}
}

func WithRetryableHTTPClientBackoff(backoff retryablehttp.Backoff) RetryableHTTPClientOption {
	return func(client *retryablehttp.Client) {
		client.Backoff = backoff
	}
}

func NewRetryableHTTPClient(opts ...RetryableHTTPClientOption) *http.Client {
	retryClient := retryablehttp.NewClient()
	// retryClient.Logger = nil
	for _, opt := range opts {
		opt(retryClient)
	}
	return retryClient.StandardClient()
}
