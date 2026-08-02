package webSockets

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	buffer     chan []byte
}

func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: connection,
		manager:    manager,
		buffer:     make(chan []byte),
	}
}

func (c *Client) readMessage() {
	defer func() {
		c.manager.removeClient(c)
	}()

	for {
		mssType, payload, err := c.connection.ReadMessage()

		if err != nil {

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}

			break
		}

		log.Println("messageType:", mssType)
		log.Println("payload:", string(payload))

	}
}

func (c *Client) writeMessage() {
	defer func() {

	}()

	for {
		select {
		case message, ok := <-c.buffer:
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}

		}
	}

}
