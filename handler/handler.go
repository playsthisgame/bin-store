package handler

import (
	"fmt"
	"log/slog"
	"plays-tcp/types"
)

var inmem = make(map[int64]*[]byte)

func Handle(cmdWrapper *types.TCPCommandWrapper) error {

	const (
		write = "WRITE"
		read  = "READ"
		list  = "LIST"
	)

	cmds := make(map[int]string)
	cmds[0] = write
	cmds[1] = read

	cmd := cmdWrapper.Command.Command

	op, ok := cmds[int(cmd)]
	if ok {
		slog.Info("operation", "op", op)
		switch op {
		case read:
			return handleRead(cmdWrapper)
		case write:
			return handleWrite(cmdWrapper)
		}
	}
	return fmt.Errorf("Unknown Operation %d", int(cmd))
}

func handleWrite(cmdWrapper *types.TCPCommandWrapper) error {
	inmem[cmdWrapper.Command.Key] = &cmdWrapper.Command.Data
	return nil
}

func handleRead(cmdWrapper *types.TCPCommandWrapper) error {
	data, ok := inmem[cmdWrapper.Command.Key]
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
