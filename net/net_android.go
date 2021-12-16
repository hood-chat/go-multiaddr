//go:build android
// +build android

package manet

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
)

// InterfaceMultiaddrs will return the addresses matching net.InterfaceAddrs
func InterfaceMultiaddrs() ([]ma.Multiaddr, error) {
	localhost, _ := FromIP(net.IPv4(127, 0, 0, 1))
	return []ma.Multiaddr{localhost}, nil
}
