package signal

type Msg struct {
	Type  int    `json:"type"`
	Code  int    `json:"code"`
	Error string `json:"error,omitzero"`
	Key   string `json:"key,omitzero"`
	Value string `json:"val,omitzero"`
}

func Start(in <-chan Msg) chan<- Msg {
	out := make(chan Msg)
	go func() {
		//
	}()

	return out
}
