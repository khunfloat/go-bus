package exam

import (
	"context"
	"go-bus/contract"
)

type ExamPort interface {
	FindAllExams(ctx context.Context) ([]contract.ExamResponse, error)
	FindExamById(ctx context.Context, request contract.ExamRequest) (*contract.ExamResponse, error)
}

type examPort struct {
	examLogic ExamLogic
}

func NewExamPort(examLogic ExamLogic) ExamPort {
	return &examPort{
		examLogic: examLogic,
	}
}

func (p *examPort) FindAllExams(ctx context.Context) ([]contract.ExamResponse, error) {
	es, err := p.examLogic.findAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]contract.ExamResponse, 0, len(es))

	for _, e := range es {
		response = append(response, contract.ExamResponse{
			ID:          e.ID,
			Name:        e.Name,
			QuestionIds: e.QuestionIds,
			Timestamp:   e.Timestamp,
		})
	}

	return response, nil
}

func (p *examPort) FindExamById(ctx context.Context, request contract.ExamRequest) (*contract.ExamResponse, error) {
	e, err := p.examLogic.findById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	response := &contract.ExamResponse{
		ID:          e.ID,
		Name:        e.Name,
		QuestionIds: e.QuestionIds,
		Timestamp:   e.Timestamp,
	}

	return response, nil
}
