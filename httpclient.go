package inet

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// HTTPClient is a simple wrapper around net/http that adds the
// Jar(http.CookieJar) method for the Client interface.
type HTTPClient struct {
	http.Client
	debug bool
}

// Debug will enable debugging/logging of Requests/Responses.
func (c *HTTPClient) Debug(enable bool) Client {
	c.debug = enable
	return c
}

// Do is a wrapper around net/http.Client.Do which allows for
// debugging of Requests/Responses.
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	var b []byte
	var e error
	var res *http.Response

	if c.debug {
		// net/http will add cookies for you, and if we add them here
		// for debugging, they end up in the request twice...
		// if c.Client.Jar != nil {
		// 	for _, cookie := range c.Client.Jar.Cookies(req.URL) {
		// 		req.AddCookie(cookie)
		// 	}
		// }

		if b, e = httputil.DumpRequestOut(req, true); e == nil {
			println(string(b))
		}
	}

	if res, e = c.Client.Do(req); e != nil {
		return res, e
	}

	if c.debug {
		if b, e = httputil.DumpResponse(res, true); e == nil {
			println(string(b))
		}
	}

	return res, nil
}

// Jar will set the cookiejar for the underlying http.Client.
func (c *HTTPClient) Jar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// Timeout will set the timeout for the underlying http.Client.
func (c *HTTPClient) Timeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// Transport will set the transport for the underlying http.Client.
func (c *HTTPClient) Transport(trans *http.Transport) Client {
	c.Client.Transport = trans
	return c
}
