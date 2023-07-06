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
	//Write fake login request packet to server
	p := rcon.NewPacket(3, "password")
	if _, err := conn.Write(p.Serialize()); err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(readPacket(conn).RequestID)
}

func readPacket(conn net.Conn) rcon.Packet {
	buf := make([]byte, 1)
	var data []byte

	// Continue reading connection stream byte-by-byte, adding bytes to data, until 2 consecutive nil bytes are reached
	// To avoid out of bounds error on expression evaluation, ensure the length of data is the minimum required for an
	// RCON packet (14  bytes)
	for len(data) < 14 || !(data[len(data)-1] == 0 && data[len(data)-2] == 0) {
		if _, err := conn.Read(buf); err != nil {
			log.Fatalf("Error reading byte: %s", err)
		}

		data = append(data, buf[0])
	}

	return rcon.Deserialize(data)
}
