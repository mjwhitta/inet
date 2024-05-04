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

// Jar will set the cookiejar for the underlying wininet.Client.
func (c *WinINetClient) Jar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// Timeout will set the timeout for the underlying wininet.Client.
func (c *WinINetClient) Timeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// Transport will set the transport for the underlying wininet.Client.
func (c *WinINetClient) Transport(trans *http.Transport) Client {
	c.Client.Transport = trans
	return c
}
