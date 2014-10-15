package gnet

import "net"

func Must(l net.Listener, e error) net.Listener {
	if e != nil {
		panic(e)
	}
	return l
}

func Listen(nt, laddr string) (net.Listener, error) {
	l, e := net.Listen(nt, laddr)
	if e != nil {
		return nil, e
	}
	return NewListener(l), nil
}
