package exam

import (
	"context"
	"time"
)

type ExamLogic interface {
	findAll(ctx context.Context) ([]ExamResponse, error)
	findById(ctx context.Context, examId string) (*ExamResponse, error)
}

type examLogic struct {
	examRepo ExamRepo
}

func NewExamLogic(examRepo ExamRepo) ExamLogic {
	return &examLogic{
		examRepo: examRepo,
	}
}

func (l *examLogic) findAll(ctx context.Context) ([]ExamResponse, error) {

	raw_data, err := l.examRepo.findAll(ctx)
	if err != nil {
		return nil, err
	}
	exams := []ExamResponse{}

	for _, exam := range raw_data {
		exams = append(exams, ExamResponse{
			ID:          exam.ExamId,
			Name:        exam.Name,
			QuestionIds: exam.QuestionIds,
			Timestamp:   time.Now().Unix(),
		})
	}

	return exams, nil
}

func (l *examLogic) findById(ctx context.Context, examId string) (*ExamResponse, error) {
	raw_data, err := l.examRepo.findById(ctx, examId)
	if err != nil {
		return nil, err
	}
	exam := &ExamResponse{
		ID:          raw_data.ExamId,
		Name:        raw_data.Name,
		QuestionIds: raw_data.QuestionIds,
		Timestamp:   time.Now().Unix(),
	}

	return exam, nil
}
