package gnet

import (
	"net"
)

func NewListener(l net.Listener) *Listener {
	return &Listener{
		Listener: l,
	}
}

type Listener struct {
	net.Listener
}

var _ net.Listener = &Listener{}
