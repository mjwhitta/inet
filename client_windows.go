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
	if e := isValidBackend(backend); e != nil {
		return e
	}

	useBackend = backend
	DefaultClient = defaultClients[useBackend]

	return nil
}

func init() {
	defaultClients = map[int]Client{HTTPBackend: &HTTPClient{}}

	if tmp, e := winhttp.NewClient(); e != nil {
		panic(e)
	} else {
		defaultClients[WinHTTPBackend] = &WinHTTPClient{*tmp}
	}

	if tmp, e := wininet.NewClient(); e != nil {
		panic(e)
	} else {
		defaultClients[WinINetBackend] = &WinINetClient{*tmp}
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
