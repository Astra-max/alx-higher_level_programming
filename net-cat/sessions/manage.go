package sessions

import (
	"net"
	"fmt"
	//"net-cat/util"
)

func NewGroup() *Group {
	g := &Group{
		Event: make(chan Event),
		ConnMembers: make(map[string]*Client),
	}

	go g.EventManager()
	return g
}

func (g *Group) NewClient(conn net.Conn) {
	conn.Write([]byte("[Enter your name]: "))
	user := make([]byte, 200)
	num,_ := conn.Read(user)

	if num == 0 {
		return
	}
	c := &Client{UserName:string(user[:num][:len(user[:num])-1]),Time:"", Conn:conn}
	g.registerClient(c, conn)
	g.Count++
	go g.clientWorker(c)
}

func (g *Group) registerClient(client *Client, conn net.Conn) {
	g.Event <- Event{Type: JOINED, Conn:client}
}

func (g *Group) clientWorker(conn *Client) {
	buf := make([]byte, 1024)

	for {
		n, err := conn.Conn.Read(buf)

		if err != nil {
			g.Event <- Event{EXITED, conn, buf[:n]}
			return
		}
		g.Event <- Event{BROADCAST, conn, buf[:n]}
	}
}

func (g *Group) BroadCast(msg string) {
	for userName, member := range g.ConnMembers {
		_,err := member.Conn.Write([]byte(msg))

		if err != nil {
			delete(g.ConnMembers, userName)
		}
	}
}

func (g *Group) EventManager() {
	
	for event := range g.Event {
		switch event.Type {
		case "JOINED":
			g.ConnMembers[event.Conn.UserName] = event.Conn
			msg := fmt.Sprintf("%s has joined our chat.\n", event.Conn.UserName)

			g.BroadCast(msg)
			break
		case "BROADCAST":
			g.BroadCast("Hello")
			break
		case "EXITED":
			msg := fmt.Sprintf("%s has left our chat.\n", event.Conn.UserName)
			g.BroadCast(msg)
			break
		default:
			break
		}
	}
}