package rcon

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

// Packet represents the standardized structure of RCON protocol packets.
type Packet struct {
	Length    int32
	RequestID int32
	Kind      int32
	Payload   []byte
}

// NewPacket returns a Packet with the user-provided Kind and Payload fields. Length and RequestID are independently
// generated.
func NewPacket(kind int32, payload string) Packet {
	return Packet{
		// Length represents the number of bytes remaining in the packet (following length). RequestID and Kind are each
		// int32 (4 bytes). The payload contains a nil-termination character (1 byte) with another nil character
		// padding following the payload (1 byte). Therefore, the length variable is 10 + payload length (in bytes).
		Length:    10 + int32(len(payload)),
		RequestID: generateID(),
		Kind:      kind,
		Payload:   []byte(payload),
	}
}

// generateID generates a random int32 variable to be used as the RequestID for the RCON packet. It is seeded with UTC
// UNIX timestamp to ensure proper variability.
func generateID() int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Int31()
}

// Serialize returns a byte slice representation of the Packet encoded per RCON protocol.
func (p *Packet) Serialize() []byte {
	// Byte buffer to store the serialized packet bytes
	var buf bytes.Buffer

	// Temporary variable to store the int32 fields as bytes
	intBytes := make([]byte, 4)

	// Convert each int32 field to little-endian and write to byte buffer
	binary.LittleEndian.PutUint32(intBytes, uint32(p.Length))
	buf.Write(intBytes)

	binary.LittleEndian.PutUint32(intBytes, uint32(p.RequestID))
	buf.Write(intBytes)

	binary.LittleEndian.PutUint32(intBytes, uint32(p.Kind))
	buf.Write(intBytes)

	// Write the payload to the byte buffer
	buf.Write(p.Payload)

	// Write two nil character bytes to the buffer to signify the end of packet
	buf.Write([]byte{0, 0})

	return buf.Bytes()
}

// TODO - create deserialize func
