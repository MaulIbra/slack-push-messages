package main

import (
	"github.com/MaulIbra/slack-push-messages/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	router := fiber.New()

	router.Get("/", func(ctx *fiber.Ctx) error {
		return nil
	})

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

	iPushNotificationUseCase := api.NewPushNotificationUseCase()
	pushNotificationDelivery := api.PushNotificationDelivery{Router: router, PushNotificationUseCase: iPushNotificationUseCase}
	pushNotificationDelivery.PushNotificationRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"

	}
	log.Fatal(router.Listen(port))
}
