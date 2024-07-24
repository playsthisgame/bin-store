package handler

import (
	"encoding/gob"
	"fmt"
	"log/slog"
	"os"
	"plays-tcp/types"
	"plays-tcp/utils"
)

var inmem = make(map[int64]*[]byte)

func Handle(cmdWrapper *types.TCPCommandWrapper) error {

	// TODO: figure out how to use iota
	const (
		write = "WRITE"
		read  = "READ"
		list  = "LIST"
		store = "STORE"
	)

	cmds := make(map[int]string)
	cmds[0] = write
	cmds[1] = read
	cmds[2] = list
	cmds[3] = store

	cmd := cmdWrapper.Command.Command

	op, ok := cmds[int(cmd)]
	if ok {
		slog.Info("operation", "op", op)
		switch op {
		case read:
			return handleRead(cmdWrapper)
		case write:
			return handleWrite(cmdWrapper)
		case store:
			return handleStore(cmdWrapper)
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
	// conn.Close() // do I want to close here?
	return fmt.Errorf("id not found")
}

func handleStore(cmdWrapper *types.TCPCommandWrapper) error {
	path := ".store" //TODO: does this need to change?
	utils.CreateDir(path)
	filename := string(cmdWrapper.Command.Data)
	encodeFile, err := os.Create(fmt.Sprintf("%s/%s.gob", path, filename)) // create the file for io
	if err != nil {
		slog.Error("error creating file for io", "error", err)
		return err
	}
	encoder := gob.NewEncoder(encodeFile)
	if err := encoder.Encode(inmem); err != nil { // Write to the file
		slog.Error("error writing to file", "error", err)
		return err
	}
	encodeFile.Close()
	return nil
}

// func handleList(cmdWrapper *types.TCPCommandWrapper) error {
//   conn := cmdWrapper.Conn
// 	keys := make([]int64, len(inmem))

// 	i := 0
// 	for k := range inmem {
// 		keys[i] = k
// 		i++
// 	}
//   conn.Writer.Writer.Write(keys)
//   return nil
// }
