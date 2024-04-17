package inet

import "net/http"

// HTTPClient is a simple wrapper around net/http that adds the
// Jar(http.CookieJar) method for the Client interface.
type HTTPClient struct {
	http.Client
}

// Jar will set the cookiejar for the underlying http.Client.
func (c *HTTPClient) Jar(jar http.CookieJar) {
	c.Client.Jar = jar
}
