package inet

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// HTTPClient is a simple wrapper around net/http that adds the
// Jar(http.CookieJar) method for the Client interface.
type HTTPClient struct {
	http.Client
	debug bool
	ua    string
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
	var cookies string
	var e error
	var res *http.Response
	var skip bool

	if (c.ua != "") && (req.Header.Get("User-Agent") == "") {
		req.Header.Set("User-Agent", c.ua)
	}

	if !c.debug {
		return c.Client.Do(req)
	}

	if b, e = httputil.DumpRequestOut(req, true); e == nil {
		if c.Client.Jar != nil {
			for _, cookie := range c.Client.Jar.Cookies(req.URL) {
				if cookies == "" {
					cookies = cookie.String()
				} else {
					cookies += "; " + cookie.String()
				}
			}
		}

		skip = cookies == ""

		for _, line := range strings.Split(string(b), "\n") {
			println(line)

			if skip {
				continue
			}

			if strings.HasPrefix(line, "Content-Length:") {
				println("Cookie: " + cookies)
				skip = true
			}
		}
	}

	if res, e = c.Client.Do(req); e != nil {
		return res, e
	}

	if b, e = httputil.DumpResponse(res, true); e == nil {
		println()
		println(string(b))
	}

	return res, nil
}

// Jar will return the Client's cookiejar.
func (c *HTTPClient) Jar() http.CookieJar {
	return c.Client.Jar
}

// SetJar will set the cookiejar for the underlying http.Client.
func (c *HTTPClient) SetJar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// SetTimeout will set the timeout for the underlying http.Client.
func (c *HTTPClient) SetTimeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// SetTransport will set the transport implementation for the
// underlying http.Client.
func (c *HTTPClient) SetTransport(trans http.RoundTripper) Client {
	c.Client.Transport = trans
	return c
}

// Timeout will return the Client's configured timeout.
func (c *HTTPClient) Timeout() time.Duration {
	return c.Client.Timeout
}

// Transport will return the Client's transport implementation.
func (c *HTTPClient) Transport() http.RoundTripper {
	return c.Client.Transport
}
