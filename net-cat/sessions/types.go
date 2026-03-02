package sessions

import "net"

const (
	JOINED = "JOINED";
	EXITED = "EXITED";
	BROADCAST = "BROADCAST"
)

type Client struct {
	UserName string
	Message string
	Time string
	Conn net.Conn
}


type Event struct {
	Type string
	Conn *Client
	Data []byte
}

type Group struct {
	Messages chan []*Client
	Event chan Event
	ConnMembers map[string]*Client
	Count int
}