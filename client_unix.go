//go:build !windows

package inet

import "github.com/mjwhitta/errors"

// Backend is used to track the preferred backend HTTP client. Only
// net/http is supported for non-Windows OS.
func Backend(backend int) error {
	if _, ok := defaultClients[backend]; !ok {
		return errors.Newf("unsupported backend %d", backend)
	}

	useBackend = HTTPBackend
	DefaultClient = defaultClients[useBackend]

	return nil
}

func init() {
	defaultClients = map[int]Client{HTTPBackend: &HTTPClient{}}

	useBackend = HTTPBackend
	DefaultClient = defaultClients[HTTPBackend]
}

// NewClient will return a new Client for the current Backend. An
// optional User-Agent can be provided.
func NewClient(ua ...string) (Client, error) {
	if len(ua) > 0 {
		return &HTTPClient{ua: ua[0]}, nil
	}

	return &HTTPClient{}, nil
}
