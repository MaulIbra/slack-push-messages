package api

import (
	"github.com/MaulIbra/slack-push-messages/api/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func UseRateLimiter(router *fiber.App) {
	router.Use(limiter.New(limiter.Config{
		Max:        1,
		Expiration: 10 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	}))
}

func Authorization(repository auth.IAuthRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		if headers["Authorization"] == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		err := repository.ClaimToken(headers["Authorization"], c.IP())
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Next()
	}
}
