package signal

import (
	"net/url"

	"github.com/gorilla/websocket"
)

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
	conn    *websocket.Conn
	in, out chan Msg
}

func NewClient(addr string) (c *Client, e error) {
	c = &Client{}
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	c.conn, _, e = websocket.DefaultDialer.Dial(u.String(), nil)
	if e != nil {
		return nil, e
	}

	return
}

func SendOffer() {

}
