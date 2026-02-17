package signal

import "net/url"

type Msg struct {
	Type  int    `json:"type"`
	Code  int    `json:"code"`
	Error string `json:"error,omitzero"`
	Key   string `json:"key,omitzero"`
	Value string `json:"val,omitzero"`
}

func start(addr string, in <-chan Msg) chan<- Msg {
	out := make(chan Msg)
	go func() {
		//
	}()

	return out
}

type Client struct {
	in, out chan Msg
}

func NewClient(addr string) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}

	return &Client{}
}

func SendOffer() {

}
