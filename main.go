package main

import (
	"github.com/k1le0/proxy/http/client"
	"github.com/k1le0/proxy/http/server"
)

func main() {

	sp := &server.Proxy{
		Name:    "",
		IP:      "127.0.0.1",
		Port:    8901,
		NetType: "tcp",
	}

	s := &server.Server{
		Proxy: sp,
	}

	cp := &client.Proxy{
		Name: "",
		IP:   "127.0.0.1",
		Port: 8901,
	}

	c := &client.Client{
		Proxy: cp,
	}

	//s.Start()

	// server

	//listener, err := net.Listen(sp.NetType, sp.IP+":"+strconv.Itoa(sp.Port))
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}

	// client

	//for {
	//	time.Sleep(time.Second)
	//	go func() {
	//		log.Println("123")
	//	}()
	//}

	//for i := 0; i < 2; i++ {
	//	time.Sleep(2 * time.Second)
	//	go func() {
	//		c.Start()
	//	}()
	//}

	go func() {
		c.Start()
	}()

	go func() {
		c.Start()
	}()

	//go func() {
	//	conn, err := net.Dial("tcp", "127.0.0.1:8901")
	//	if err != nil {
	//		log.Println(err.Error())
	//	}
	//	defer func(conn net.Conn) {
	//		err := conn.Close()
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}(conn)
	//	for j := 0; j < 5; j++ {
	//		time.Sleep(time.Second)
	//
	//		_, err := conn.Write([]byte("hello"))
	//		if err != nil {
	//			log.Println(err)
	//			continue
	//		}
	//	}
	//	isClose := "CLOSE"
	//	_, err = conn.Write([]byte(isClose))
	//	if err != nil {
	//		log.Println(err.Error())
	//		return
	//	}
	//}()

	//go func() {
	//	time.Sleep(10 * time.Second)
	//	conn, err := net.Dial("tcp", "127.0.0.1:8901")
	//	if err != nil {
	//		log.Println(err.Error())
	//	}
	//	defer func(conn net.Conn) {
	//		err := conn.Close()
	//		if err != nil {
	//
	//		}
	//	}(conn)
	//	for j := 0; j < 5; j++ {
	//		time.Sleep(time.Second)
	//
	//		_, err := conn.Write([]byte("hi"))
	//		if err != nil {
	//			log.Println(err)
	//			continue
	//		}
	//	}
	//	isClose := "CLOSE"
	//	_, err = conn.Write([]byte(isClose))
	//	if err != nil {
	//		log.Println(err.Error())
	//		return
	//	}
	//}()

	//server mes stream
	s.Start()

	//for {
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	go func() {
	//		for {
	//			buf := make([]byte, 512)
	//			length, err := conn.Read(buf)
	//			if err != nil {
	//				log.Println(err)
	//				continue
	//			}
	//
	//			msg := string(buf[:length])
	//			if "CLOSE" == msg {
	//				break
	//			} else if "HEARTBEAT" == msg {
	//				log.Println("remote: " + conn.RemoteAddr().String())
	//			} else {
	//				log.Println("message: " + msg)
	//			}
	//		}
	//		err := conn.Close()
	//		if err != nil {
	//			return
	//		}
	//	}()
	//}
}
