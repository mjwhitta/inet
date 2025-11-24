package inet

import (
	"io"
	"net/http"
	"net/url"
)

// Get will make a GET request using the DefaultClient.
func Get(uri string) (res *http.Response, e error) {
	if res, e = DefaultClient.Get(uri); e != nil {
		return nil, e //nolint:wrapcheck // Intentionally not wrapping
	}

	return res, nil
}

// Head will make a HEAD request using the DefaultClient.
func Head(uri string) (res *http.Response, e error) {
	if res, e = DefaultClient.Head(uri); e != nil {
		return nil, e //nolint:wrapcheck // Intentionally not wrapping
	}

	return res, nil
}

// Post will make a POST request using the DefaultClient.
func Post(
	uri string,
	contentType string,
	body io.Reader,
) (res *http.Response, e error) {
	if res, e = DefaultClient.Post(uri, contentType, body); e != nil {
		return nil, e //nolint:wrapcheck // Intentionally not wrapping
	}

	return res, nil
}

// PostForm will make a POST request using the DefaultClient.
func PostForm(
	uri string,
	form url.Values,
) (res *http.Response, e error) {
	if res, e = DefaultClient.PostForm(uri, form); e != nil {
		return nil, e //nolint:wrapcheck // Intentionally not wrapping
	}

	return res, nil
}
