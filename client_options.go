package tapd

import "net/http"

type ClientOption func(*Client) error

// WithBaseURL sets the baseURL for the client
func WithBaseURL(urlStr string) ClientOption {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithBasicAuth sets the username and password for the client
func WithBasicAuth(username, password string) ClientOption {
	return func(c *Client) error {
		c.username = username
		c.password = password
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}
