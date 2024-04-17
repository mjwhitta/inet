//go:build windows

package inet

import (
	"net/http"

	"github.com/mjwhitta/win/winhttp"
)

// WinHTTPClient is a simple wrapper around winhttp that adds the
// Jar(http.CookieJar) method for the Client interface.
type WinHTTPClient struct {
	winhttp.Client
}

// Jar will set the cookiejar for the underlying winhttp.Client.
func (c *WinHTTPClient) Jar(jar http.CookieJar) {
	c.Client.Jar = jar
}
