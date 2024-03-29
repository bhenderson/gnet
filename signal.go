package gnet

import (
	"net"
	"os"
	"os/signal"
)

func Signal(l net.Listener, sigs ...os.Signal) {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, sigs...)

	for sig := range ch {
		for _, si := range sigs {
			if sig == si {
				if gl, ok := l.(*Listener); ok {
					gl.closec <- struct{}{}
				}
				break // inner loop
			}
		}
	}
}
