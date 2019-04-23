// Package generic adapts the lego exec DNS
// provider for Caddy. Importing this package plugs it in.
package acmedns

import (
	"errors"

	"github.com/mholt/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/acmedns"
)

func init() {
	caddytls.RegisterDNSProvider("acmedns", NewDNSProvider)
}

// NewDNSProvider returns a new exec DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use program to run from environment variable EXEC_PATH
// len(1): credentials[0] = program to run
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return acmedns.NewDNSProvider()
	/*case 1:
		config := acmedns.NewDefaultConfig()
		config.Program = credentials[0]
		return acmedns.NewDNSProviderConfig(config)
	*/
	default:
		return nil, errors.New("invalid credentials length")
	}
}