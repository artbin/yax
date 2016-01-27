package netx

import (
	"net"
	"os"
	"strings"

	"github.com/ArtemKulyabin/yax/netx/npipex"
)

func Listen(network, laddr string) (net.Listener, error) {
	if network == "npipe" {
		if !strings.HasPrefix(laddr, `\\.\pipe\`) {
			laddr = `\\.\pipe\` + laddr
		}
		return npipex.Listen(laddr)
	} else {
		return net.Listen(network, laddr)
	}
}

func Dial(network, address string) (net.Conn, error) {
	if network == "npipe" {
		if !strings.HasPrefix(address, `\\.\pipe\`) {
			address = `\\.\pipe\` + address
		}
		return npipex.Dial(address)
	} else {
		return net.Dial(network, address)
	}
}

func FileListener(f *os.File) (net.Listener, error) {
	l, err := npipex.FileListener(f)
	if err != nil {
		return net.FileListener(f)
	}
	return l, nil
}
