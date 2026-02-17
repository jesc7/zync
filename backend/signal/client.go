package signal

import (
	"net/url"
	"sync"

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
	mu      sync.RWMutex
}

func NewClient(addr string) (c *Client, e error) {
	c = &Client{
		in:  make(chan Msg),
		out: make(chan Msg),
	}
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	c.conn, _, e = websocket.DefaultDialer.Dial(u.String(), nil)
	if e != nil {
		return nil, e
	}
	return
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	close(c.in)
	close(c.out)
	c.conn.Close()
}

func SendOffer() {

}
