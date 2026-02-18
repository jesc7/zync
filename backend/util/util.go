package util

import (
	"encoding/json"
	"errors"
	"os"
)

func IsFileExists(filename string) bool {
	if _, e := os.Stat(filename); errors.Is(e, os.ErrNotExist) {
		return false
	}
	return true
}

func JSONEncode(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func JSONDecode(s string, dst any) error {
	return json.Unmarshal([]byte(s), dst)
}
