package gnet

import (
	"errors"
	"net"
	"sync"
)

var StoppedError = errors.New("Server was stopped.")

func NewListener(l net.Listener) *Listener {
	gl := &Listener{
		Listener: l,
		closec:   make(chan struct{}),
		acceptc:  make(chan *Conn),
	}
	go gl.accept()
	return gl
}

type Listener struct {
	wg      sync.WaitGroup
	closec  chan struct{}
	acceptc chan *Conn
	net.Listener
}

var _ net.Listener = &Listener{}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case <-l.closec:
		l.wg.Wait()
	case c := <-l.acceptc:
		l.wg.Add(1)
		c.wg = &l.wg
		return c, c.e
	}
	return nil, StoppedError
}

func (l *Listener) accept() {
	for {
		l.acceptc <- newConn(l.Listener.Accept())
	}
}
