package proxy

import (
	"net"
	"net/url"
)

type Proxy interface {
	Start()
	Stop()
	Handle(conn net.Conn)
	HeartBeat(conn net.Conn)
	Message(conn net.Conn)
}

type Server struct {
}

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Handle() {

}

type ProxyCreator func(url *url.URL) (Proxy, error)

var (
	proxyMap = make(map[string]ProxyCreator)
)

func RegisterProxy(name string, pc ProxyCreator) {
	proxyMap[name] = pc
}
