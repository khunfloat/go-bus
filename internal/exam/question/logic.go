package question

import (
	"context"
	"time"
)

// interface
type QuestionLogic interface {
	findAll(ctx context.Context) ([]QuestionResponse, error)
	findById(ctx context.Context, questionId string) (*QuestionResponse, error)
}

// struct
type questionLogic struct {
	questionRepo QuestionRepo
}

// constructor
func NewQuestionLogic(questionRepo QuestionRepo) QuestionLogic {
	return &questionLogic{
		questionRepo: questionRepo,
	}
}

// findAll
func (l *questionLogic) findAll(ctx context.Context) ([]QuestionResponse, error) {
	rawData, err := l.questionRepo.findAll(ctx)
	if err != nil {
		return nil, err
	}

	questions := []QuestionResponse{}
	for _, q := range rawData {
		questions = append(questions, QuestionResponse{
			ID:        q.QuestionId,
			Question:  q.Question,
			Answer:    q.Answer,
			Timestamp: time.Now().Unix(),
		})
	}

	return questions, nil
}

// findById
func (l *questionLogic) findById(ctx context.Context, questionId string) (*QuestionResponse, error) {
	rawData, err := l.questionRepo.findById(ctx, questionId)
	if err != nil {
		return nil, err
	}

	question := &QuestionResponse{
		ID:        rawData.QuestionId,
		Question:  rawData.Question,
		Answer:    rawData.Answer,
		Timestamp: time.Now().Unix(),
	}

	return question, nil
}
