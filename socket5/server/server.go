package server

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// 维护一个clientMap
var (
	serverMap = make(map[string]ServConn)
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

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err.Error())
				break
			}

			tcpListener := listener.(*net.TCPListener)
			err = tcpListener.SetDeadline(time.Now().Add(time.Second * 5))
			if err != nil {
				return
			}

			dealConn := &ServConn{
				Conn: conn,
				Uid:  rand.Int31(),
			}
			serverMap[dealConn.Conn.RemoteAddr().String()] = *dealConn
		}
	}()

	for _, v := range serverMap {
		//conn, err := listener.Accept()
		//if err != nil {
		//	log.Println(err.Error())
		//	continue
		//}
		//dealConn := &ServConn{
		//	Conn: conn,
		//	Uid:  rand.Int31(),
		//}
		//serverMap[dealConn.Conn.RemoteAddr().String()] = *dealConn

		go s.Handle(v.Conn)
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
		}

		log.Println("message: " + string(buf[:length]))

		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err.Error())
			continue
		}
	}
}

func (s *Server) HeartBeat(conn net.Conn) {
}

func (s *Server) Message(conn net.Conn) {

}
