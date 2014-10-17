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
	// This effectively turns the Accept() method into a channel so we can
	// select on it. Otherwise, if we were blocking on Accept() we would miss
	// the close channel.
	//
	// This also has the added benefit of not calling Accept() until we're
	// ready, as opposed to in the initializer which would be somewhat
	// unexpected.
	for {
		l.acceptc <- newConn(l.Listener.Accept())
	}
}
