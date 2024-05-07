//go:build !windows

package inet

// Backend is used to track the preferred backend HTTP client. Only
// net/http is supported for non-Windows OS.
func Backend(backend int) error {
	if e := isValidBackend(backend); e != nil {
		return e
	}

	useBackend = HTTPBackend
	DefaultClient = defaultClients[useBackend]

	return nil
}

func init() {
	useBackend = HTTPBackend
	defaultClients = map[int]Client{HTTPBackend: &HTTPClient{}}
	DefaultClient = defaultClients[HTTPBackend]
}

// NewClient will return a new Client for the current Backend. An
// optional User-Agent can be provided for Windows backends only.
// User-Agent will still need to be specified for requests.
func NewClient(ua ...string) (Client, error) {
	if len(ua) > 0 {
		return &HTTPClient{ua: ua[0]}, nil
	}

	return &HTTPClient{}, nil
}
