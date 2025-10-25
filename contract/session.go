package contract

import "context"

type SessionContract interface {
	FindSessionById(ctx context.Context, request SessionRequest) (*SessionResponse, error)
	CreateSession(ctx context.Context, request CreateSessionRequest) (*SessionResponse, error)
}

type SessionResponse struct {
	ID          string
	UserId      string
	ExamId      string
	QuestionIds []string
	Timestamp   int64
}

type CreateSessionRequest struct {
	UserId string
	ExamId string
}

type SessionRequest struct {
	ID string
}
