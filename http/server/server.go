package server

import (
	"log"
	"net"
	"strconv"
)

type Server struct {
	Proxy *Proxy
}

type Proxy struct {
	Name    string
	IP      string
	Port    int
	NetType string
}

func (s *Server) Start() {
	listener, err := net.Listen(s.Proxy.NetType, s.Proxy.IP+":"+strconv.Itoa(s.Proxy.Port))
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go s.Handle(conn)
		//buf := make([]byte, 512)
		//
		//length, err := conn.Read(buf)
		//if err != nil {
		//	log.Println(err.Error())
		//	continue
		//}
		//log.Println("message: " + string(buf[:length]))
		//
		//_, err = conn.Write(buf)
		//if err != nil {
		//	log.Println(err)
		//}
	}
}

func (s *Server) Stop() {}

func (s *Server) Handle(conn net.Conn) {
	buf := make([]byte, 512)

	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println("message: " + string(buf[:length]))

		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}
