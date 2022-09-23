package main

import (
	"fmt"
	"github.com/MaulIbra/slack-push-messages/api"
	"github.com/MaulIbra/slack-push-messages/api/auth"
	"github.com/MaulIbra/slack-push-messages/api/push_notification"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	router := fiber.New()

	mySigningKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	expiredToken, _ := strconv.Atoi(os.Getenv("EXPIRED_TIME_TOKEN"))

	api.UseRateLimiter(router)

	iAuthRepo := auth.NewAuthRepository(token, mySigningKey, expiredToken)
	iAuthUseCase := auth.NewAuthUseCase(iAuthRepo)
	authDelivery := auth.DeliveryAuth{Router: router, AuthUseCase: iAuthUseCase}
	authDelivery.AuthDeliveryRoutes()

	router.Use(api.Authorization(iAuthRepo))

	router.Get("/", func(ctx *fiber.Ctx) error {
		return nil
	})

	iPushNotificationUseCase := push_notification.NewPushNotificationUseCase()
	pushNotificationDelivery := push_notification.PushNotificationDelivery{Router: router, PushNotificationUseCase: iPushNotificationUseCase}
	pushNotificationDelivery.PushNotificationRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"

	}
	log.Fatal(router.Listen(fmt.Sprintf(":%v", port)))
}
