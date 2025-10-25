package exam

import (
	"go-bus/contract"
	"go-bus/internal/exam/exam"
	"go-bus/internal/exam/question"
	"go-bus/pkg/bus"

	"github.com/gofiber/fiber/v2"
)

func InitExamService(app fiber.Router, bus *bus.RequestBus) {

	// init db
	db := initDb("exam-dsn")

	// exam
	examRepo := exam.NewExamRepo(db)
	examLogic := exam.NewExamLogic(examRepo)
	examPort := exam.NewExamPort(examLogic)

	// question
	questionRepo := question.NewQuestionRepo(db)
	questionLogic := question.NewQuestionLogic(questionRepo)
	questionPort := question.NewQuestionPort(questionLogic)

	// service port
	servicePort := newServicePort(examPort, questionPort)
	bus.RegisterContract("Exam", (*contract.ExamContract)(nil), servicePort)

	// Api
	examApi := app.Group("/exam")
	exam.RegisterHandlers(examApi, examLogic)

	questionApi := app.Group("/question")
	question.RegisterHandlers(questionApi, questionLogic)
}

func initDb(dsn string) string {
	db := "init " + dsn
	return db
}

type servicePort struct {
	exam.ExamPort
	question.QuestionPort
}

func newServicePort(examPort exam.ExamPort, questionPort question.QuestionPort) contract.ExamContract {
	return &servicePort{
		examPort,
		questionPort,
	}
}
