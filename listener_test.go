package gnet

type tlistener struct{}

var (
	_     net.Listener = &tlistener{}
	tconn              = &net.TCPConn{}
	taddr              = &net.TCPAddr{}
)

func (tl *tlistener) Accept() (net.Conn, error) {
	return tconn, nil
}

func (tl *tlistener) Close() error {
	return nil
}

func (tl *tlistener) Addr() net.Addr {
	return taddr
}
