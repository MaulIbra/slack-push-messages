package main

import (
	"github.com/MaulIbra/slack-push-messages/api"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	router := fiber.New()
	iPushNotificationUseCase := api.NewPushNotificationUseCase()
	pushNotificationDelivery := api.PushNotificationDelivery{Router: router, PushNotificationUseCase: iPushNotificationUseCase}
	pushNotificationDelivery.PushNotificationRoutes()
	log.Fatal(router.Listen(":3000"))
}
