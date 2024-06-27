//go:build !windows

package inet

import "github.com/mjwhitta/errors"

func isValidBackend(backend int) error {
	switch backend {
	case HTTPBackend:
	default:
		return errors.Newf("unsupported backend %d", backend)
	}

	return nil
}
