// +build !arm

package providers

import "time"

func providerTimeout() time.Duration {
	return 15 * time.Second
}
