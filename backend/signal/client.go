package signal

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"

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
	conn         *websocket.Conn
	in, out, ans chan Msg
	mu           sync.RWMutex
	ctx          context.Context
}

func NewClient(ctx context.Context, addr string) (c *Client, e error) {
	c = &Client{
		in:  make(chan Msg),
		out: make(chan Msg),
		ans: make(chan Msg),
	}
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	c.conn, _, e = websocket.DefaultDialer.Dial(u.String(), nil)
	if e != nil {
		return nil, e
	}

	var cancel context.CancelFunc
	c.ctx, cancel = context.WithCancel(ctx)
	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case m := <-c.out:
				if e = c.conn.WriteJSON(m); e != nil {
					return
				}
			}
		}
	}()

	go func() {
		defer cancel()

		var m Msg
		for {
			if e = c.conn.ReadJSON(&m); e != nil || ctx.Err() != nil {
				return
			}
			switch m.Type {
			case MT_PING:
				c.out <- Msg{Type: MT_PONG}
			case MT_RECEIVEANSWER: //к клиенту кто-то хочет подключиться
				c.ans <- Msg{Type: MT_PONG}
			default:
				c.in <- m
			}
		}
	}()
	return
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	close(c.in)
	close(c.out)
	close(c.ans)
	c.conn.Close()
}

func (c *Client) read(timeout time.Duration) (m Msg, e error) {
	select {
	case <-c.ctx.Done():
		return Msg{}, c.ctx.Err()
	case <-time.After(timeout):
		return Msg{}, errors.New("Read timeout")
	case m = <-c.in:
		return
	}
}

func (c *Client) SendOffer(offer string) (key, password string, e error) {
	if c.ctx.Err() != nil {
		return "", "", c.ctx.Err()
	}

	c.out <- Msg{
		Type:  MT_SENDOFFER,
		Value: offer,
	}

	answer, e := c.read(10 * time.Second)
	if e != nil {
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
