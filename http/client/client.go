package client

import (
	"log"
	"net"
	"strconv"
)

type Client struct {
	Proxy *Proxy
}

type Proxy struct {
	Name string
	IP   string
	Port int
}

func (c *Client) Start() {
	conn, err := net.Dial("tcp", c.Proxy.IP+":"+strconv.Itoa(c.Proxy.Port))
	if err != nil {
		log.Println(err.Error())
	}

	go c.Handle(conn)
}

func (c *Client) Stop() {}

func (c *Client) Handle(conn net.Conn) {
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	//isClose := "CLOSE"
	//_, err := conn.Write([]byte(isClose))
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
}
