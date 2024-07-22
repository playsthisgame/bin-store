package types

import (
	"encoding/binary"
	"fmt"
)

var VERSION byte = 0
var HEADER_SIZE = 12

type TCPCommand struct {
	Command byte
	Data    []byte
	Key     int64
}

// verion(1) + command(1) + length(2) + key(8) + data(?)
func (t *TCPCommand) MarshalBinary() (data []byte, err error) {
	length := uint16(len(t.Data))                      // get the length of the data
	lengthData := make([]byte, 2)                      // make a byte array of size 2
	keyData := make([]byte, 8)                         // make a byte array of size 8 to hold the key
	binary.BigEndian.PutUint16(lengthData, length)     // put the length in the lengthData byte array
	binary.BigEndian.PutUint64(keyData, uint64(t.Key)) // put the key in the keyData byte array

	b := make([]byte, 0, 1+1+2+8+length) // make a byte array of size 0 and a capacity of the HEADER_SIZE + the length of data
	b = append(b, VERSION)               // add the version to the byte array
	b = append(b, t.Command)             // append the Command
	b = append(b, lengthData...)         // append the length of the data, since its 2 bytes it will spread them into the return byte array
	b = append(b, keyData...)            // append the length of the data, since its 2 bytes it will spread them into the return byte array
	return append(b, t.Data...), nil     // append the data itself into the byte array and a nil for the error
}

func (t *TCPCommand) UnmarshalBinary(bytes []byte) error {
	if bytes[0] != VERSION { // if the first byte is not the correct version then return an error
		return fmt.Errorf("version mismatch %d != %d", bytes[0], VERSION)
	}

	length := int(binary.BigEndian.Uint16(bytes[2:])) // get the length at the second index, Uint16 --> 16 bits which is 2 bytes
	end := HEADER_SIZE + length                       // calculate the end

	if len(bytes) < end {
		return fmt.Errorf("not enough data to parse packet: got %d expected %d", len(bytes), HEADER_SIZE+length)
	}

	command := bytes[1]
	key := bytes[4:12]
	data := bytes[HEADER_SIZE:end]

	t.Command = command

	t.Key = int64(binary.BigEndian.Uint64(key))
	t.Data = data

	return nil
}

type TCPCommandWrapper struct {
	Conn    *Connection
	Command *TCPCommand
}
