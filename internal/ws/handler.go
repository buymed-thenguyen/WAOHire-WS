package ws

import (
	"backend-ws/internal/room"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(rm *room.RoomManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionCode := c.Param("session_code")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		rm.AddClient(sessionCode, conn)
		defer rm.RemoveClient(sessionCode, conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}
}
