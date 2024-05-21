//go:build windows

package inet

import (
	"net/http"
	"time"

	"github.com/mjwhitta/win/wininet"
)

// WinINetClient is a simple wrapper around wininet that adds the
// Jar(http.CookieJar) method for the Client interface.
type WinINetClient struct {
	wininet.Client
}

// Debug will enable debugging/logging of Requests/Responses.
func (c *WinINetClient) Debug(enable bool) Client {
	c.Client.Debug = enable
	return c
}

// Jar will return the wininet.Client's cookiejar.
func (c *WinINetClient) Jar() http.CookieJar {
	return c.Client.Jar
}

// SetJar will set the cookiejar for the underlying wininet.Client.
func (c *WinINetClient) SetJar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// SetTimeout will set the timeout for the underlying wininet.Client.
func (c *WinINetClient) SetTimeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// SetTransport will set the transport implementation for the
// underlying wininet.Client.
func (c *WinINetClient) SetTransport(trans http.RoundTripper) Client {
	c.Client.Transport = trans
	return c
}

// Timeout will return the wininet.Client's configured timeout.
func (c *WinINetClient) Timeout() time.Duration {
	return c.Client.Timeout
}

// Transport will return the wininet.Client's transport
// implementation.
func (c *WinINetClient) Transport() http.RoundTripper {
	return c.Client.Transport
}
