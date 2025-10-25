package main

import (
	"go-bus/internal/exam"
	"go-bus/internal/gateway"
	"go-bus/internal/session"
	"go-bus/pkg/bus"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	bus := bus.NewRequestBus()

	router := gateway.ApiGateway(app)

	exam.InitExamService(router, bus)
	session.InitSessionService(router, bus)

	log.Fatal(app.Listen(":3001"))
}
