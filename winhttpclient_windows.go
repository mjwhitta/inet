//go:build windows

package inet

import (
	"net/http"
	"time"

	"github.com/mjwhitta/win/winhttp"
)

// WinHTTPClient is a simple wrapper around winhttp that adds the
// Jar(http.CookieJar) method for the Client interface.
type WinHTTPClient struct {
	winhttp.Client
}

// Debug will enable debugging/logging of Requests/Responses.
func (c *WinHTTPClient) Debug(enable bool) Client {
	c.Client.Debug = enable
	return c
}

// Jar will return the winhttp.Client's cookiejar.
func (c *WinHTTPClient) Jar() http.CookieJar {
	return c.Client.Jar
}

// SetJar will set the cookiejar for the underlying winhttp.Client.
func (c *WinHTTPClient) SetJar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// SetTimeout will set the timeout for the underlying winhttp.Client.
func (c *WinHTTPClient) SetTimeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// SetTransport will set the transport implementation for the
// underlying winhttp.Client.
func (c *WinHTTPClient) SetTransport(trans http.RoundTripper) Client {
	c.Client.Transport = trans
	return c
}

// Timeout will return the winhttp.Client's configured timeout.
func (c *WinHTTPClient) Timeout() time.Duration {
	return c.Client.Timeout
}

// Transport will return the winhttp.Client's transport
// implementation.
func (c *WinHTTPClient) Transport() http.RoundTripper {
	return c.Client.Transport
}
