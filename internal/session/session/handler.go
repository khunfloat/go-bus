package session

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(r fiber.Router, l SessionLogic) {
	r.Get("/:id", func(c *fiber.Ctx) error {
		questionId := c.Params("id")
		ctx := c.Context()
		question, err := l.findById(ctx, questionId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(question)
	})

	r.Post("/start", func(c *fiber.Ctx) error {
		var req SessionRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		ctx := c.Context()
		session, err := l.create(ctx, req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(session)
	})
}
