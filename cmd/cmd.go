package cmd

import (
	"backend-ws/config"
	"backend-ws/internal/grpc"
	"backend-ws/internal/room"
	"backend-ws/internal/ws"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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

	rm := room.NewRoomManager()
	go grpc.StartGRPCServer(rm, cfg.GrpcPort)

	r := gin.Default()
	r.GET("/ws/:session_code", ws.HandleWS(rm))

	port := os.Getenv("PORT") // fallback for railway deployment
	if port == "" {
		port = cfg.Port
	}
	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
