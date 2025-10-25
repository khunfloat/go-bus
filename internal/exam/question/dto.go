package question

type QuestionResponse struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	Timestamp int64  `json:"timestamp"`
}
