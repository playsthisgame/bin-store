package handler

import (
	"fmt"
	"log/slog"
	"plays-tcp/types"
)

var inmem = make(map[string]*[]byte)

func Handle(cmdWrapper *types.TCPCommandWrapper) error {

	const (
		write = "WRITE"
		read  = "READ"
	)

	cmds := make(map[int]string)
	cmds[0] = write
	cmds[1] = read

	cmd := cmdWrapper.Command.Command
	data := cmdWrapper.Command.Data

	op, ok := cmds[int(cmd)]
	if ok {
		slog.Info("operation", "op", op)
		switch op {
		case read:
			return handleRead(cmdWrapper)
		case write:
			return handleWrite(&data)
		}
	}
	return fmt.Errorf("Unknown Operation %d", int(cmd))
}

func handleWrite(data *[]byte) error {
	inmem["test"] = data
	return nil
}

func handleRead(cmdWrapper *types.TCPCommandWrapper) error {
	data, ok := inmem["test"]
	conn := cmdWrapper.Conn
	if ok {
		_, err := conn.Writer.Writer.Write(*data)
		if err != nil {
			return err
		}
		return nil
	}
	conn.Close()
	return fmt.Errorf("id not found")
}
