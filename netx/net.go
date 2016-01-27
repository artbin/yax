// +build !windows

package netx

import (
	"net"
	"os"
)

func Listen(network, laddr string) (net.Listener, error) {
	return net.Listen(network, laddr)
}

func Dial(network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}

func FileListener(f *os.File) (ln net.Listener, err error) {
	return net.FileListener(f)
}
