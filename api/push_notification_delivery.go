package api

import "github.com/gofiber/fiber/v2"

type PushNotificationDelivery struct {
	Router                  *fiber.App
	PushNotificationUseCase IPushNotificationUseCase
}

func (pnd *PushNotificationDelivery) PushNotificationRoutes() {
	pnd.Router.Post("/send-message", pnd.SendMessage)
}

func (pnd *PushNotificationDelivery) SendMessage(c *fiber.Ctx) error {

	pushNotification := &PushNotification{}

	if err := c.BodyParser(pushNotification); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := pushNotification.ValidateStruct()
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	err := pnd.PushNotificationUseCase.SendMessage(pushNotification)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "message sent successfully",
	})
}
