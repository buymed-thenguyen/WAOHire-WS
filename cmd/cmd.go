package cmd

import (
	"backend-ws/config"
	"backend-ws/internal/domain"
	"backend-ws/internal/grpc"
	"backend-ws/internal/handler"
	"backend-ws/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	if cfg == nil {
		panic("empty config")
	}
	fmt.Println("✅ Config loaded")

	// logger
	logger.InitLogger(cfg.Logger.Path)
	fmt.Println("✅ Initiated logger")

	rm := domain.NewRoomManager()

	// grpc
	go grpc.StartGRPCServer(rm, cfg.GrpcPort)

	r := gin.Default()

	// ws
	r.GET("/ws/:session_code", handler.HandleWS(rm))

	// http
	r.POST("/ws/user-joined", handler.UserJoined)
	r.POST("/ws/user-leaved", handler.UserLeaved)
	r.POST("/ws/user-answered", handler.UserAnswered)
	r.POST("/ws/start-session", handler.StartSession)

	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
