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
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	PostForm(string, url.Values) (*http.Response, error)

	// These functions are unique to this module.
	Debug(enable bool) Client
	Jar() http.CookieJar
	SetJar(jar http.CookieJar) Client
	SetTimeout(timeout time.Duration) Client
	SetTransport(trans http.RoundTripper) Client
	Timeout() time.Duration
	Transport() http.RoundTripper
}
