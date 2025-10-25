package session

type SessionResponse struct {
	ID          string   `json:"id"`
	UserId      string   `json:"userId"`
	ExamId      string   `json:"examId"`
	QuestionIds []string `json:"questionIds"`
	Timestamp   int64    `json:"timestamp"`
}

type SessionRequest struct {
	UserId string `json:"userId"`
	ExamId string `json:"examId"`
}
