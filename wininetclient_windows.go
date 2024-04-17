//go:build windows

package inet

import (
	"net/http"

	"github.com/mjwhitta/win/wininet"
)

// WinINetClient is a simple wrapper around wininet that adds the
// Jar(http.CookieJar) method for the Client interface.
type WinINetClient struct {
	wininet.Client
}

// Jar will set the cookiejar for the underlying wininet.Client.
func (c *WinINetClient) Jar(jar http.CookieJar) {
	c.Client.Jar = jar
}
