package room

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type RoomManager struct {
	rooms map[string]map[*websocket.Conn]bool
	mu    sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]map[*websocket.Conn]bool),
	}
}

func (r *RoomManager) AddClient(sessionCode string, conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.rooms[sessionCode] == nil {
		r.rooms[sessionCode] = make(map[*websocket.Conn]bool)
	}
	r.rooms[sessionCode][conn] = true
}

func (r *RoomManager) RemoveClient(sessionCode string, conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if clients, ok := r.rooms[sessionCode]; ok {
		delete(clients, conn)
		if len(clients) == 0 {
			delete(r.rooms, sessionCode)
		}
	}
}

func (r *RoomManager) Broadcast(key string, message []byte) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if clients, ok := r.rooms[key]; ok {
		for conn := range clients {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Send message failed: %v", err)
			}
		}
	}
}
