package exam

import (
	"context"
	"go-bus/internal/exam/entity"
)

type ExamRepo interface {
	findAll(ctx context.Context) ([]entity.Exam, error)
	findById(ctx context.Context, examId string) (*entity.Exam, error)
}

type examRepo struct {
	db string
}

func NewExamRepo(db string) ExamRepo {
	return &examRepo{
		db: db,
	}
}

func (r *examRepo) findAll(ctx context.Context) ([]entity.Exam, error) {
	exams := []entity.Exam{
		{
			ExamId:      "exam-1",
			Name:        "Math Exam",
			QuestionIds: []string{"q1", "q2", "q3"},
		},
		{
			ExamId:      "exam-2",
			Name:        "Science Exam",
			QuestionIds: []string{"q4", "q5", "q6"},
		},
	}
	return exams, nil
}

func (r *examRepo) findById(ctx context.Context, examId string) (*entity.Exam, error) {
	exam := &entity.Exam{
		ExamId:      examId,
		Name:        "Sample Exam",
		QuestionIds: []string{"q1", "q2", "q10"},
	}
	return exam, nil
}
