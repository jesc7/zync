package backend

import "encoding/json"

func JSONEncode(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func JSONDecode(s string, dst any) error {
	return json.Unmarshal([]byte(s), dst)
}
