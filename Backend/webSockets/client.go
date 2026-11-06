package webSockets

import (
	"encoding/json"
	"fmt"
	"log"
	"matchMaker/model"
	"time"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	buffer     chan []byte
}

const (
	pongWait   = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: connection,
		manager:    manager,
		buffer:     make(chan []byte, 256),
	}
}

func (c *Client) ReadMessage() {
	defer func() {
		c.manager.removeClient(c)
	}()

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		fmt.Println("Error setting read deadline:", err)
		return
	}

	// set read Limit
	c.connection.SetReadLimit(512)
	c.connection.SetPongHandler(c.PongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}
		var receivedMsg model.MsgType
		error := json.Unmarshal(payload, &receivedMsg)
		if error != nil {
			fmt.Println("Error unmarshalling message:", error.Error())
			return
		}
		log.Println("payload:", string(payload))

	}
}

func (c *Client) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.buffer:
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("write error:", err)
				return
			}
		case <-ticker.C:
			c.connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) WriteMessage(message []byte) {
	c.buffer <- message
}

func (c *Client) PongHandler(pong string) error {
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
