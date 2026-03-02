package main

import (
	"os"
	"net"
	"log"
	"fmt"
	"net-cat/util"
	"net-cat/sessions"
)

func main() {
	err, port := util.GetPort()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("server listening on port %v\n", port)
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Println("Error creating listener", err.Error())
		return
	}

	logo, err := os.ReadFile("icon")

	if err != nil {
		log.Println("Failed to load chat Icon")
		return
	}

	for {
		conn, err := listener.Accept()
		newGroup := sessions.NewGroup()

		if newGroup.Count > 3 {
			log.Println("Group full please wait for member to exit")
		}

		conn.Write(logo)
		go newGroup.NewClient(conn)
		if err != nil {
			log.Println("Failed to connect to client", err.Error())
			return
		}
	}
}