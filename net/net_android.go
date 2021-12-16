//go:build android
// +build android

package manet

type androidInterface struct {}

// @FIXME(gfanton): on android sdk30, syscall from `net.InterfaceAddrs()` are restricted, manually return localhost on android
func (androidInterface) InterfaceAddrs() ([]net.Addr, error) { return net.IPv4(127, 0, 0, 1), nil }

func getNetDriver() NetDriver {
	return &androidInterface{}
}
