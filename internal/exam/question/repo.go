package question

import (
	"context"
	"go-bus/internal/exam/entity"
)

type QuestionRepo interface {
	findAll(ctx context.Context) ([]entity.Question, error)
	findById(ctx context.Context, QuestionId string) (*entity.Question, error)
}

type questionRepo struct {
	db string
}

func NewQuestionRepo(db string) QuestionRepo {
	return &questionRepo{
		db: db,
	}
}

func (r *questionRepo) findAll(ctx context.Context) ([]entity.Question, error) {
	questions := []entity.Question{
		{
			QuestionId: "q1",
			Question:   "q1 question",
			Answer:     "q1 answer",
		},
		{
			QuestionId: "q2",
			Question:   "q2 question",
			Answer:     "q2 answer",
		},
		{
			QuestionId: "q3",
			Question:   "q3 question",
			Answer:     "q3 answer",
		},
	}

	return questions, nil
}

func (r *questionRepo) findById(ctx context.Context, QuestionId string) (*entity.Question, error) {
	question := &entity.Question{
		QuestionId: QuestionId,
		Question:   "sample question",
		Answer:     "sample answer",
	}

	return question, nil
}
