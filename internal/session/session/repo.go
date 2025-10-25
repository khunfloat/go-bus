package session

import (
	"context"
	"go-bus/internal/session/entity"
)

type SessionRepo interface {
	findById(ctx context.Context, sessionId string) (*entity.Session, error)
	create(ctx context.Context, session entity.Session) (*entity.Session, error)
}

type sessionRepo struct {
	db string
}

func NewSessionRepo(db string) SessionRepo {
	return &sessionRepo{
		db: db,
	}
}

func (r *sessionRepo) findById(ctx context.Context, sessionId string) (*entity.Session, error) {
	s := entity.Session{
		SessionId:   sessionId,
		UserId:      "user123",
		ExamId:      "exam456",
		QuestionIds: []string{"q4", "q5", "q6"},
	}

	return &s, nil
}

func (r *sessionRepo) create(ctx context.Context, session entity.Session) (*entity.Session, error) {
	s := entity.Session{
		SessionId:   "newSessionId",
		UserId:      session.UserId,
		ExamId:      session.ExamId,
		QuestionIds: session.QuestionIds,
	}

	return &s, nil
}
