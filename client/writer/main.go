package main

import (
	"crypto/rand"
	"io"
	"log/slog"
	"net"
	"time"

	"github.com/playsthisgame/bin-store/types"
)

// an example of a client
func main() {

	// TODO: for now the key will be the current unix time which is an int64 (8 bytes), a problem with this is that I can run into a race condition if two writes happen within the same ms
	key := time.Now().Unix()
	// k := make([]byte, 8)
	// binary.BigEndian.PutUint64(k, uint64(key))

	file := make([]byte, 100)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		slog.Error("Error creating file:", "error", err)
	}

	cmd := &types.TCPCommand{
		Command: 0,
		Data:    file,
		Key:     key,
	}

	data, err := cmd.MarshalBinary()
	if err != nil {
		slog.Error("Error marshalling data:", "error", err)
	}

	// dial server
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		slog.Error("Error dialing server:", "error", err)
	}

	defer conn.Close()

	_, err = conn.Write(data)
	if err != nil {
		slog.Error("Error dialing server:", "error", err)
	}

	slog.Info("written %d bytes over the network", "bytes", len(data))
}
