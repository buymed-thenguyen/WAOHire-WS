package cmd

import (
	"backend-ws/config"
	"backend-ws/internal/domain"
	"backend-ws/internal/grpc"
	"backend-ws/internal/handler"
	"backend-ws/utils/logger"
	"fmt"
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

	r := handler.SetupRouter(rm)
	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
