package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nhutHao02/social-network-chat-service/internal/domain/model"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"go.uber.org/zap"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Socket struct {
	//          map[roomID]map[userID]*websocket.Conn
	connections         map[string]map[string]*websocket.Conn
	mu                  sync.RWMutex
	recentMsgConnection map[string]*websocket.Conn
}

func NewSocket() *Socket {
	return &Socket{
		connections:         make(map[string]map[string]*websocket.Conn),
		recentMsgConnection: make(map[string]*websocket.Conn),
	}
}

// Add connection
func (s *Socket) AddConnection(roomID string, userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// check room exist, if unexist -> create
	if _, exists := s.connections[roomID]; !exists {
		s.connections[roomID] = make(map[string]*websocket.Conn)
	}
	s.connections[roomID][userID] = conn
}

// Remove connection
func (s *Socket) RemoveConnection(roomID string, userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.connections[roomID], userID)
	conn.Close()
}

// Broadcast message to all connections
func (s *Socket) Broadcast(roomID string, userID string, message model.OutgoingMessageWSRes) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, conn := range s.connections[roomID] {
		// if uid != userID {
		if err := conn.WriteJSON(message); err != nil {
			logger.Error("Socket-Broadcast: Error sending message", zap.Error(err))
			s.RemoveConnection(roomID, userID, conn)
		}
		// }
	}
}

// Add connection
func (s *Socket) AddRecentMgsConnection(userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.recentMsgConnection[userID] = conn
}

// Remove connection
func (s *Socket) RemoveRecentMgsConnection(userID string, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.recentMsgConnection, userID)
	conn.Close()
}

// Broadcast message to all connections
func (s *Socket) BroadcastRecentMgs(userID string, message model.OutgoingMessageWSRes) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for uid, conn := range s.recentMsgConnection {
		if uid == userID {
			if err := conn.WriteJSON(message); err != nil {
				logger.Error("Socket-Broadcast: Error sending recent message", zap.Error(err))
				s.RemoveRecentMgsConnection(userID, conn)
			}
		}
	}
}
