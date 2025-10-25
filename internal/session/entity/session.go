package entity

type Session struct {
	SessionId   string
	UserId      string
	ExamId      string
	QuestionIds []string
}
