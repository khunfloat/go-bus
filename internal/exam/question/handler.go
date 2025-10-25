package question

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(r fiber.Router, l QuestionLogic) {
	r.Get("/", func(c *fiber.Ctx) error {
		ctx := c.Context()
		questions, err := l.findAll(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(questions)
	})

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
}
