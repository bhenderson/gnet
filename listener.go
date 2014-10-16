package gnet

import (
	"net"
	"sync"
)

func NewListener(l net.Listener) *Listener {
	return &Listener{
		Listener: l,
	}
}

type Listener struct {
	wg sync.WaitGroup
	net.Listener
}

var _ net.Listener = &Listener{}

func (l *Listener) Accept() (net.Conn, error) {
	// do something
	l.wg.Add(1)
	c, e := l.Listener.Accept()
	return &Conn{l.wg, c}, e
}

type Conn struct {
	wg sync.WaitGroup
	net.Conn
}

func (c *Conn) Close() error {
	c.wg.Done()
	return c.Conn.Close()
}
