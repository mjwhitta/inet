package inet

// Supported backends
const (
	HTTPBackend = iota // Default on non-Windows
	WinHTTPBackend
	WinINetBackend // Default on Windows
)

var (
	// DefaultClient points to the default Client for the current
	// backend.
	DefaultClient  Client
	defaultClients map[int]Client
	useBackend     int
)

// Version is the package version
const Version string = "0.6.0"
