package subdir

import "net"

func Fn() {
	net.ParseIP("foobar")
}
