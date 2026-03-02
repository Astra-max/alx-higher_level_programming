package main

import (
	"net"
	"log"
	"net-cat/util"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	buf := make([]byte, 1024)

	if err != nil {
		log.Println("Failed to establish connections to remote server", err.Error())
		return
	}

	user := util.UserInput("Enter your name: ")

	for {
		data := util.UserInput("Leave server a note: ")

		if data == "quit" {
			return
		}
		_, err := conn.Write([]byte(user + ": " + data + "\n"))

		if err != nil{
			log.Println("Error writing to connected server", err.Error())
		}

		_, err = conn.Read(buf)

		if err != nil {
			log.Println("Failed to read from server", err.Error())
			return
		}

		fmt.Println(string(buf))
	}
}