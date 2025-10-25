package contract

import "context"

type ExamContract interface {
	FindAllExams(ctx context.Context) ([]ExamResponse, error)
	FindExamById(ctx context.Context, request ExamRequest) (*ExamResponse, error)
	FindAllQuestions(ctx context.Context) ([]QuestionResponse, error)
	FindQuestionById(ctx context.Context, request QuestionRequest) (*QuestionResponse, error)
}

type ExamRequest struct {
	ID string
}

type QuestionRequest struct {
	ID string
}

type ExamResponse struct {
	ID          string
	Name        string
	QuestionIds []string
	Timestamp   int64
}

type QuestionResponse struct {
	ID        string
	Question  string
	Answer    string
	Timestamp int64
}
