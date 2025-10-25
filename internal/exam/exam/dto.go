package exam

type ExamResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	QuestionIds []string `json:"questionIds"`
	Timestamp   int64    `json:"timestamp"`
}
