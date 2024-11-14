package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// server program to handle concurrent client connections.

	listener, error := net.Listen("tcp", "localhost:8000")
	if error != nil {
		log.Fatal(error)
	}

	// check if there is no infinite loop then what is happening.
	// 1. with handdleConn without go routine, you can't be able to connect with multiple clients.
	// 2. after adding go routine on handleConnection as soon as connection will start with client it will away
	// to avoid this for some time you can use time.sleep and then let see.
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//  time.Sleep(10*time.Second)
		go handleConn(connection)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
