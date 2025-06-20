package request

type BroadcastRequest struct {
	UserID      uint   `json:"user_id"`
	SessionCode string `json:"session_code"`
	QuizID      uint   `json:"quiz_id"`
}
