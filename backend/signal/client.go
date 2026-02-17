package signal

import (
	"errors"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type MessageType int

const (
	MT_NOANSWER      MessageType = iota - 1 //не отправлять ответ
	MT_SENDOFFER                            //клиент1 отправил offer
	MT_SENDANSWER                           //клиент2 отправил answer
	MT_RECEIVEANSWER                        //клиенту1 отправили answer клиента2
	MT_CONNECT                              //клиент1 уведомляет об установлении соединения
	MT_DISCONNECT                           //клиент1 уведомляет о разрыве соединения
	MT_PING                                 //ping
	MT_PONG                                 //pong
)

type Msg struct {
	Type  MessageType `json:"type"`
	Code  int         `json:"code"`
	Key   string      `json:"key,omitzero"`
	Value string      `json:"val,omitzero"`
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

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	return
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	close(c.in)
	close(c.out)
	c.conn.Close()
}

func (c *Client) SendOffer(offer string) (key, password string, e error) {
	if e = c.conn.WriteJSON(Msg{
		Type:  MT_SENDOFFER,
		Value: offer,
	}); e != nil {
		return
	}

	var answer Msg
	if e = c.conn.ReadJSON(&answer); e != nil {
		return
	}
	if answer.Type != MT_SENDOFFER {
		return "", "", errors.New("Wrong answer type")
	}
	if answer.Code == -1 {
		return "", "", errors.New(answer.Value)
	}

	sl := strings.Split(answer.Key, "@")
	if len(sl) < 2 {
		return "", "", errors.New("Wrong key value")
	}
	return sl[0], sl[1], nil
}
