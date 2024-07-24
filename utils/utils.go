package utils

import (
	"errors"
	"log/slog"
	"os"
)

func CreateDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			slog.Error("error creating store", "error", err)
		}
	}
}
