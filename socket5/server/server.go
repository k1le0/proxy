package server

import (
	"github.com/k1le0/proxy/socket5/client"
	"log"
	"math/rand"
	"net"
	"strconv"
)

// 维护一个clientMap
var (
	clientMap = make(map[string]client.CliConn)
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

type ServConn struct {
	Conn net.Conn
	Uid  int32
}

func (s *Server) Start() {
	listener, err := net.Listen(s.Proxy.NetType, s.Proxy.IP+":"+strconv.Itoa(s.Proxy.Port))
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		dealConn := &ServConn{
			Conn: conn,
			Uid:  rand.Int31(),
		}
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go s.Handle(dealConn.Conn)
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

func (s *Server) Stop() {

}

func (s *Server) Handle(conn net.Conn) {
	buf := make([]byte, 512)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		msg := string(buf[:length])
		if "CLOSE" == msg {
			break
		} else if "HEARTBEAT" == msg {
			HeartBeat(conn)
		} else {
			log.Println("message: " + string(buf[:length]))
		}

		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func HeartBeat(conn net.Conn) {
	ipStr := conn.RemoteAddr().String()
	c := &client.CliConn{Conn: conn, Uid: rand.Int31()}
	clientMap[ipStr] = *c
}

func Message(conn net.Conn) {

}
