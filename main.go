package main

import (
	"log/slog"
	"os"

	"github.com/playsthisgame/bin-store/handler"
	"github.com/playsthisgame/bin-store/tcp"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	server, err := tcp.NewTCPServer(3000)
	if err != nil {
		slog.Error("Error creating new TCP server:", "error", err)
	}
	defer server.Close()
	go server.Start()

	for {
		cmd := <-server.FromSockets
		handler.Handle(&cmd)
	}
}
