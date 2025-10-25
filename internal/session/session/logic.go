package session

import (
	"context"
	"go-bus/contract"
	"go-bus/internal/session/entity"
	"go-bus/pkg/bus"
	"time"
)

// interface
type SessionLogic interface {
	findById(ctx context.Context, sessionId string) (*SessionResponse, error)
	create(ctx context.Context, session SessionRequest) (*SessionResponse, error)
}

// struct
type sessionLogic struct {
	sessionRepo SessionRepo
	bus         *bus.RequestBus
}

// constructor
func NewSessionLogic(sessionRepo SessionRepo, bus *bus.RequestBus) SessionLogic {
	return &sessionLogic{
		sessionRepo: sessionRepo,
		bus:         bus,
	}
}

// findAll
func (l *sessionLogic) findById(ctx context.Context, sessionId string) (*SessionResponse, error) {
	rawData, err := l.sessionRepo.findById(ctx, sessionId)
	if err != nil {
		return nil, err
	}

	s := &SessionResponse{
		ID:          rawData.SessionId,
		UserId:      rawData.UserId,
		ExamId:      rawData.ExamId,
		QuestionIds: rawData.QuestionIds,
		Timestamp:   time.Now().Unix(),
	}

	return s, nil
}

// findById
func (l *sessionLogic) create(ctx context.Context, session SessionRequest) (*SessionResponse, error) {

	examResult, err := l.bus.CallContract(
		ctx,
		"Exam.FindExamById",
		contract.ExamRequest{ID: session.ExamId},
	)

	exam := examResult.(*contract.ExamResponse)

	if err != nil {
		return nil, err
	}

	rawData, err := l.sessionRepo.create(ctx, entity.Session{
		UserId:      session.UserId,
		ExamId:      exam.ID,
		QuestionIds: exam.QuestionIds,
	})

	if err != nil {
		return nil, err
	}

	s := &SessionResponse{
		ID:          rawData.SessionId,
		UserId:      rawData.UserId,
		ExamId:      rawData.ExamId,
		QuestionIds: rawData.QuestionIds,
		Timestamp:   time.Now().Unix(),
	}

	return s, nil
}
