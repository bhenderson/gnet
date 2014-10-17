package gnet

import (
	"net"
	"sync"
)

func newConn(c net.Conn, e error) *Conn {
	return &Conn{
		nil, c, e,
	}
}

type Conn struct {
	wg *sync.WaitGroup
	net.Conn
	e error
}

func (c *Conn) Close() error {
	c.wg.Done()
	return c.Conn.Close()
}

var _ net.Conn = &Conn{}
