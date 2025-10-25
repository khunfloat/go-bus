package session

import (
	"go-bus/contract"
	"go-bus/internal/session/session"
	"go-bus/pkg/bus"

	"github.com/gofiber/fiber/v2"
)

func InitSessionService(app fiber.Router, bus *bus.RequestBus) {

	// init db
	db := initDb("session-dsn")

	// session
	sessionRepo := session.NewSessionRepo(db)
	sessionLogic := session.NewSessionLogic(sessionRepo, bus)
	sessionPort := session.NewSessionPort(sessionLogic)

	// service port
	servicePort := newServicePort(sessionPort)
	bus.RegisterContract("Session", (*contract.SessionContract)(nil), servicePort)

	// Api
	sessionApi := app.Group("/session")
	session.RegisterHandlers(sessionApi, sessionLogic)
}

func initDb(dsn string) string {
	db := "init " + dsn
	return db
}

type servicePort struct {
	session.SessionPort
}

func newServicePort(sessionPort session.SessionPort) contract.SessionContract {
	return &servicePort{
		sessionPort,
	}
}
