package session

import (
	"context"
	"go-bus/contract"
)

type SessionPort interface {
	FindSessionById(ctx context.Context, request contract.SessionRequest) (*contract.SessionResponse, error)
	CreateSession(ctx context.Context, request contract.CreateSessionRequest) (*contract.SessionResponse, error)
}

type sessionPort struct {
	sessionLogic SessionLogic
}

func NewSessionPort(sessionLogic SessionLogic) SessionPort {
	return &sessionPort{
		sessionLogic: sessionLogic,
	}
}

func (p *sessionPort) FindSessionById(ctx context.Context, request contract.SessionRequest) (*contract.SessionResponse, error) {
	s, err := p.sessionLogic.findById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	response := &contract.SessionResponse{
		ID:          s.ID,
		UserId:      s.UserId,
		ExamId:      s.ExamId,
		QuestionIds: s.QuestionIds,
		Timestamp:   s.Timestamp,
	}

	return response, nil
}

func (p *sessionPort) CreateSession(ctx context.Context, request contract.CreateSessionRequest) (*contract.SessionResponse, error) {
	s, err := p.sessionLogic.create(ctx, SessionRequest{
		UserId: request.UserId,
		ExamId: request.ExamId,
	})
	if err != nil {
		return nil, err
	}

	response := &contract.SessionResponse{
		ID:          s.ID,
		UserId:      s.UserId,
		ExamId:      s.ExamId,
		QuestionIds: s.QuestionIds,
		Timestamp:   s.Timestamp,
	}

	return response, nil
}
