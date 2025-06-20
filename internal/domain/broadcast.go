package domain

import (
	"backend-ws/model/request"
	"backend-ws/model/response"
	"encoding/json"
)

func UserAnswered(req *request.BroadcastRequest) *response.DefaultResponse {
	msg := map[string]interface{}{
		"event":        "user_answered",
		"user_id":      req.UserID,
		"session_code": req.SessionCode,
	}
	jsonData, _ := json.Marshal(msg)
	rm.Broadcast(req.SessionCode, jsonData)
	return &response.DefaultResponse{Message: "ok"}
}

func UserJoined(req *request.BroadcastRequest) *response.DefaultResponse {
	msg := map[string]interface{}{
		"event": "user_joined",
	}
	jsonData, _ := json.Marshal(msg)
	rm.Broadcast(req.SessionCode, jsonData)
	return &response.DefaultResponse{Message: "ok"}
}

func UserLeaved(req *request.BroadcastRequest) *response.DefaultResponse {
	msg := map[string]interface{}{
		"event": "user_leaved",
	}
	jsonData, _ := json.Marshal(msg)
	rm.Broadcast(req.SessionCode, jsonData)
	return &response.DefaultResponse{Message: "ok"}
}

func StartSession(req *request.BroadcastRequest) *response.DefaultResponse {
	msg := map[string]interface{}{
		"event":        "start_session",
		"quiz_id":      req.QuizID,
		"session_code": req.SessionCode,
	}
	jsonData, _ := json.Marshal(msg)
	rm.Broadcast(req.SessionCode, jsonData)
	return &response.DefaultResponse{Message: "ok"}
}
