package grpc

import (
	"backend-ws/internal/domain"
	"backend-ws/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

type Server struct {
	proto.UnimplementedBroadcasterServer
	rm *domain.RoomManager
}

func NewServer(rm *domain.RoomManager) *Server {
	return &Server{rm: rm}
}

func (s *Server) BroadcastUserAnswered(c context.Context, req *proto.UserAnsweredRequest) (*proto.BroadcastResponse, error) {
	msg := map[string]interface{}{
		"event":        "user_answered",
		"user_id":      req.UserId,
		"session_code": req.SessionCode,
	}
	jsonData, _ := json.Marshal(msg)
	s.rm.Broadcast(req.SessionCode, jsonData)
	return &proto.BroadcastResponse{
		Success: true,
	}, nil
}

func (s *Server) BroadcastUserJoined(c context.Context, req *proto.UserJoinedRequest) (*proto.BroadcastResponse, error) {
	msg := map[string]interface{}{
		"event": "user_joined",
	}
	jsonData, _ := json.Marshal(msg)
	s.rm.Broadcast(req.SessionCode, jsonData)
	return &proto.BroadcastResponse{
		Success: true,
	}, nil
}

func (s *Server) BroadcastUserLeaved(c context.Context, req *proto.UserLeavedRequest) (*proto.BroadcastResponse, error) {
	msg := map[string]interface{}{
		"event": "user_leaved",
	}
	jsonData, _ := json.Marshal(msg)
	s.rm.Broadcast(req.SessionCode, jsonData)
	return &proto.BroadcastResponse{
		Success: true,
	}, nil
}

func (s *Server) BroadcastStartSession(c context.Context, req *proto.StartSessionRequest) (*proto.BroadcastResponse, error) {
	msg := map[string]interface{}{
		"event":        "start_session",
		"quiz_id":      req.QuizId,
		"session_code": req.SessionCode,
	}
	jsonData, _ := json.Marshal(msg)
	s.rm.Broadcast(req.SessionCode, jsonData)
	return &proto.BroadcastResponse{
		Success: true,
	}, nil
}

func StartGRPCServer(rm *domain.RoomManager, port string) {
	envPort := os.Getenv("GRPC_PORT") // fallback port for railway deployment
	if envPort == "" {
		envPort = port
	}

	lis, err := net.Listen("tcp", ":"+envPort)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterBroadcasterServer(grpcServer, NewServer(rm))
	reflection.Register(grpcServer)
	fmt.Println("gRPC server started on :" + port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
