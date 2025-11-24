package inet

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

// Client is an interface defining required functions for an HTTP
// client.
type Client interface {
	// The following functions mirror net/http functions.
	Do(req *http.Request) (*http.Response, error)
	Get(uri string) (*http.Response, error)
	Head(uri string) (*http.Response, error)
	Post(
		uri string,
		contentType string,
		body io.Reader,
	) (*http.Response, error)
	PostForm(uri string, form url.Values) (*http.Response, error)

	// These functions are unique to this module.
	Debug(enable bool) Client
	Jar() http.CookieJar
	SetJar(jar http.CookieJar) Client
	SetTimeout(timeout time.Duration) Client
	SetTransport(trans http.RoundTripper) Client
	Timeout() time.Duration
	Transport() http.RoundTripper
}
