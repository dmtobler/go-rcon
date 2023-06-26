package main

import (
	"log"
	"net"
)

func main() {

	const (
		HOST = "tobler.games"
		PORT = "25575"
		TYPE = "tcp"
	)

	// Create TCP connection
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error establishing connection to %s:%s", HOST, PORT)
	}
	defer conn.Close()

}
