//go:build windows

package inet

import (
	"github.com/mjwhitta/errors"
	"github.com/mjwhitta/win/winhttp"
	"github.com/mjwhitta/win/wininet"
)

// Backend is used to track the preferred backend HTTP client. Windows
// allows for net/http, WinHTTP, and WinINet. The default is WinINet.
func Backend(backend int) error {
	if _, ok := defaultClients[backend]; !ok {
		return errors.Newf("unsupported backend %d", backend)
	}

	useBackend = backend
	DefaultClient = defaultClients[useBackend]

	return nil
}

func init() {
	var c1 *winhttp.Client
	var c2 *wininet.Client
	var e error

	if c1, e = winhttp.NewClient(); e != nil {
		panic(errors.Newf("failed to create client: %w", e))
	}

	if c2, e = wininet.NewClient(); e != nil {
		panic(errors.Newf("failed to create client: %w", e))
	}

	defaultClients = map[int]Client{
		HTTPBackend:    &HTTPClient{},
		WinHTTPBackend: &WinHTTPClient{*c1},
		WinINetBackend: &WinINetClient{*c2},
	}

	useBackend = WinINetBackend
	DefaultClient = defaultClients[useBackend]
}

// NewClient will return a new Client for the current Backend. An
// optional User-Agent can be provided.
func NewClient(ua ...string) (Client, error) {
	var c1 *winhttp.Client
	var c2 *wininet.Client
	var e error

	switch useBackend {
	case HTTPBackend:
		if len(ua) > 0 {
			return &HTTPClient{ua: ua[0]}, nil
		}

		return &HTTPClient{}, nil
	case WinHTTPBackend:
		if c1, e = winhttp.NewClient(ua...); e != nil {
			return nil, errors.Newf("failed to create client: %w", e)
		}

		return &WinHTTPClient{*c1}, nil
	default:
		if c2, e = wininet.NewClient(ua...); e != nil {
			return nil, errors.Newf("failed to create client: %w", e)
		}

		return &WinINetClient{*c2}, nil
	}
}
