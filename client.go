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
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
	Head(url string) (*http.Response, error)
	Jar(jar http.CookieJar)
	Post(
		url string, contentType string, body io.Reader,
	) (*http.Response, error)
	PostForm(url string, data url.Values) (*http.Response, error)
	Timeout(timeout time.Duration)
}
