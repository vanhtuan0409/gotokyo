package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

type ClientOpt struct {
	Addr    string
	BotName string

	DebugRead  bool
	DebugWrite bool
}

type Client struct {
	conn *websocket.Conn
	opt  *ClientOpt

	sendLimiter *rate.Limiter
}

func NewClient(opt ClientOpt) (*Client, error) {
	if opt.BotName == "" || opt.Addr == "" {
		return nil, errors.New("Invalid client config")
	}

	c := Client{
		opt:         &opt,
		sendLimiter: rate.NewLimiter(ClientSendRate, 1),
	}
	if err := c.connect(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Client) connect() error {
	wsAddr := fmt.Sprintf("%s/socket?key=%s&name=%s", c.opt.Addr, c.opt.BotName, c.opt.BotName)
	conn, _, err := websocket.DefaultDialer.Dial(wsAddr, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) ReadEvent(ctx context.Context) (*Event, error) {
	_, r, err := c.conn.NextReader()
	if err != nil {
		return nil, err
	}
	if c.opt.DebugRead {
		r = io.TeeReader(r, os.Stdout)
	}

	var e Event
	if err := json.NewDecoder(r).Decode(&e); err != nil {
		return nil, err
	}
	return &e, nil
}

type jsonCommand struct {
	Type string      `json:"e"`
	Data interface{} `json:"data,omitempty"`
}

func (c *Client) SendCommand(ctx context.Context, command Command) error {
	c.sendLimiter.Wait(ctx)

	w, err := c.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	defer w.Close()

	var wt io.Writer = w
	if c.opt.DebugWrite {
		wt = io.MultiWriter(w, os.Stdout)
	}

	msg := jsonCommand{
		Type: command.Name(),
		Data: command.Data(),
	}
	return json.NewEncoder(wt).Encode(&msg)
}

func (c *Client) Close() error {
	c.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	)
	return c.conn.Close()
}
