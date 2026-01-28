package pkg

import (
	"errors"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c fiber.Ctx, err error) error {

		code := fiber.StatusInternalServerError
		msg  := "internal server error"

		var fe *fiber.Error
		if errors.As(err, &fe) {
			code = fe.Code
			msg  = fe.Message
		}

		log.Printf(
			"[ERROR] %s %s\nerr=%v\nstack=%s",
			c.Method(),
			c.Path(),
			err,
			debug.Stack(),
		)

		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"message": msg,
			"code":    code,
		})
	}
}
