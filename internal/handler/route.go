package handler

import (
	"backend-ws/internal/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(rm *domain.RoomManager) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// ws
	r.GET("/ws/:session_code", HandleWS(rm))

	// http
	r.POST("/ws/user-joined", UserJoined)
	r.POST("/ws/user-leaved", UserLeaved)
	r.POST("/ws/user-answered", UserAnswered)
	r.POST("/ws/start-session", StartSession)

	// html
	r.Static("/template", "./template")
	r.StaticFile("/", "./template/index.html")
	r.StaticFile("/log-out", "./template/logout.html")
	r.StaticFile("/quizzes", "./template/quizzes.html")
	r.StaticFile("/session", "./template/session.html")
	r.StaticFile("/question", "./template/question.html")
	r.StaticFile("/leaderboard", "./template/leaderboard.html")

	return r
}
