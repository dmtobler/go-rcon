package main

import (
	"log"
	"net"
)

type Packet struct {
	Length    int32
	RequestID int32
	Kind      int32
	Payload   string
}

func (p Packet) toString(b []byte) string {
	return string(b[:])
}

//func (p Packet) generateID() int32 {
//
//}

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
