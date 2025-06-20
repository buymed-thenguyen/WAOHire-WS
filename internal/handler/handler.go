package handler

import (
	"backend-ws/internal/domain"
	"backend-ws/model/request"
	"backend-ws/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWS(rm *domain.RoomManager) gin.HandlerFunc {
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

func UserJoined(c *gin.Context) {
	var req *request.BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	domain.UserJoined(req)
}

func UserLeaved(c *gin.Context) {
	var req *request.BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	domain.UserLeaved(req)
}

func UserAnswered(c *gin.Context) {
	var req *request.BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	domain.UserAnswered(req)
}

func StartSession(c *gin.Context) {
	var req *request.BroadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.BadRequest(c, err.Error())
		return
	}

	domain.StartSession(req)
}
