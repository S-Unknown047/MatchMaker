package webSockets

import (
	"encoding/json"
	"matchMaker/model"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type EventHandler func(event model.Event, c *Client) error
type Manager struct {
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	m := &Manager{
		clients: make(ClientList),
	}

	return m
}

func (m *Manager) Init(w http.ResponseWriter, r *http.Request) {

	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)
	go client.ReadMessage()
	go client.writeMessage()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}

}
