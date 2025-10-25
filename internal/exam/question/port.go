package question

import (
	"context"
	"go-bus/contract"
)

type QuestionPort interface {
	FindAllQuestions(ctx context.Context) ([]contract.QuestionResponse, error)
	FindQuestionById(ctx context.Context, request contract.QuestionRequest) (*contract.QuestionResponse, error)
}

type questionPort struct {
	quesetionLogic QuestionLogic
}

func NewQuestionPort(quesetionLogic QuestionLogic) QuestionPort {
	return &questionPort{
		quesetionLogic: quesetionLogic,
	}
}

func (p *questionPort) FindAllQuestions(ctx context.Context) ([]contract.QuestionResponse, error) {
	qs, err := p.quesetionLogic.findAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]contract.QuestionResponse, 0, len(qs))

	for _, q := range qs {
		response = append(response, contract.QuestionResponse{
			ID:        q.ID,
			Question:  q.Question,
			Answer:    q.Answer,
			Timestamp: q.Timestamp,
		})
	}

	return response, nil
}

func (p *questionPort) FindQuestionById(ctx context.Context, request contract.QuestionRequest) (*contract.QuestionResponse, error) {
	q, err := p.quesetionLogic.findById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	response := &contract.QuestionResponse{
		ID:        q.ID,
		Question:  q.Question,
		Answer:    q.Answer,
		Timestamp: q.Timestamp,
	}

	return response, nil
}
