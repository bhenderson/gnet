package gnet

import (
	"net"
	"time"
)

func Must(l net.Listener, e error) net.Listener {
	if e != nil {
		panic(e)
	}
	return l
}

func Listen(nt, laddr string) (l net.Listener, e error) {
	l, e = net.Listen(nt, laddr)
	if e != nil {
		return
	}
	// http.ListenAndServe()
	if tl, ok := l.(*net.TCPListener); ok {
		l = &tcpKeepAliveListener{tl}
	}
	return NewListener(l), nil
}

// copied from net/http/server.go
// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
