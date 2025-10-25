package exam

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(r fiber.Router, l ExamLogic) {
	r.Get("/", func(c *fiber.Ctx) error {
		ctx := c.Context()
		exams, err := l.findAll(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(exams)
	})

	r.Get("/:id", func(c *fiber.Ctx) error {
		examId := c.Params("id")
		ctx := c.Context()
		exam, err := l.findById(ctx, examId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(exam)
	})
}
