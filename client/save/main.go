package main

import (
	"log/slog"
	"net"
	"plays-tcp/types"
)

// an example of a client
func main() {
	filename := "test"
	cmd := &types.TCPCommand{
		Command: 2,
		Data:    []byte(filename),
		Key:     0,
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

	// received := make([]byte, 1024)
	// _, err = conn.Read(received)
	// if err != nil {
	// 	println("Read data failed:", err.Error())
	// 	os.Exit(1)
	// }
	// slog.Info("Received data", "data", string(received))
}
