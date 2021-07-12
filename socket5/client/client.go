package client

import (
	"log"
	"net"
	"strconv"
	"time"
)

type Client struct {
	Proxy *Proxy
}

type Proxy struct {
	Name string
	IP   string
	Port int
}

type CliConn struct {
	Conn net.Conn
	Uid  int32
}

func (c *Client) Start() {
	conn, err := net.Dial("tcp", c.Proxy.IP+":"+strconv.Itoa(c.Proxy.Port))
	if err != nil {
		log.Println(err.Error())
	}

	go c.HeartBeat(conn)
	c.Handle(conn)
}

func (c *Client) Stop() {}

func (c *Client) Handle(conn net.Conn) {
	for {
		time.Sleep(2 * time.Second)
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			log.Println(err)
			continue
		}
	}
	isClose := "CLOSE"
	_, err := conn.Write([]byte(isClose))
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (c *Client) HeartBeat(conn net.Conn) {
	for {
		time.Sleep(5 * time.Second)
		_, err := conn.Write([]byte("HEARTBEAT"))
		if err != nil {
			return
		}
	}
}

func (c *Client) Message(conn net.Conn) {

}
