package fcoin

import (
	"github.com/gorilla/websocket"
	"time"
)

func (c *Client) InitWS() error {
	conn, _, err := websocket.DefaultDialer.Dial(WSBaseUrl, nil)
	if err != nil {
		return err
	}
	c.WS = conn

	// discard Hello from server
	var rsp WSHello
	if err := c.WS.ReadJSON(&rsp); err != nil {
		c.WS = nil
		return err
	}
	return nil
}

func (c *Client) WSPing() error {
	args := WSArgs{
		Cmd: "ping",
	}
	t := time.Now().Unix() * 1000
	args.Args = append(args.Args, t)
	if err := c.WS.WriteJSON(args); err != nil {
		return err
	}
	return nil
}

func (c *Client) action(action, id string, topics ...interface{}) error {
	args := WSArgs{
		Cmd: action,
	}
	for _, v := range topics {
		args.Args = append(args.Args, v)
	}

	if id != "" {
		args.ID = id
	}
	if err := c.WS.WriteJSON(args); err != nil {
		return err
	}
	return nil
}

// Add new subscription
func (c *Client) WSSubscribe(id string, topics ...interface{}) error {
	return c.action("sub", id, topics...)
}

// Un-subscription
func (c *Client) WSUnsubscribe(id string, topics ...interface{}) error {
	return c.action("unsub", id, topics...)
}

// Request once
func (c *Client) WSReq(id string, args ...interface{}) error {
	return c.action("req", id, args...)
}
