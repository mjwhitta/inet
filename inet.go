package inet

import (
	"io"
	"net/http"
	"net/url"
)

// Get will make a GET request using the DefaultClient.
func Get(url string) (*http.Response, error) {
	return DefaultClient.Get(url)
}

// Head will make a HEAD request using the DefaultClient.
func Head(url string) (*http.Response, error) {
	return DefaultClient.Head(url)
}

// Post will make a POST request using the DefaultClient.
func Post(
	url string, contentType string, body io.Reader,
) (*http.Response, error) {
	return DefaultClient.Post(url, contentType, body)
}

// PostForm will make a POST request using the DefaultClient.
func PostForm(url string, data url.Values) (*http.Response, error) {
	return DefaultClient.PostForm(url, data)
}
