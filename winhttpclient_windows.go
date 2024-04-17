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

// Jar will set the cookiejar for the underlying winhttp.Client.
func (c *WinHTTPClient) Jar(jar http.CookieJar) Client {
	c.Client.Jar = jar
	return c
}

// Timeout will set the timeout for the underlying winhttp.Client.
func (c *WinHTTPClient) Timeout(timeout time.Duration) Client {
	c.Client.Timeout = timeout
	return c
}

// Transport will set the transport for the underlying winhttp.Client.
func (c *WinHTTPClient) Transport(trans *http.Transport) Client {
	c.Client.Transport = trans
	return c
}
