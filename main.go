package main

import (
	"fmt"
	"github.com/dmtobler/go-rcon/rcon"
	"log"
	"net"
)

func main() {
	// TODO - refactor to future connection file
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

	// TODO - remove below

	// Write fake login request packet to server
	b := make([]byte, 64)
	p := rcon.NewPacket(3, "password")
	if _, err := conn.Write(p.Serialize()); err != nil {
		log.Fatalf("%s", err)
	}

	if _, err := conn.Read(b); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(b)
}
