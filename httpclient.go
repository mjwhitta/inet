package inet

import (
	"net/http"
	"time"
)

// HTTPClient is a simple wrapper around net/http that adds the
// Jar(http.CookieJar) method for the Client interface.
type HTTPClient struct {
	http.Client
}

// Jar will set the cookiejar for the underlying http.Client.
func (c *HTTPClient) Jar(jar http.CookieJar) {
	c.Client.Jar = jar
}

// Timeout will set the timeout for the underlying http.Client.
func (c *HTTPClient) Timeout(timeout time.Duration) {
	c.Client.Timeout = timeout
}
