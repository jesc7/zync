package util

import (
	"errors"
	"os"
)

func IsFileExists(filename string) bool {
	if _, e := os.Stat(filename); errors.Is(e, os.ErrNotExist) {
		return false
	}
	return true
}
