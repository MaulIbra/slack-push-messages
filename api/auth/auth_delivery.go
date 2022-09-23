package auth

import "github.com/gofiber/fiber/v2"

type DeliveryAuth struct {
	Router      *fiber.App
	AuthUseCase IAuthUseCase
}

func (da *DeliveryAuth) AuthDeliveryRoutes() {
	da.Router.Post("/token", da.GenerateToken)
}

func (da *DeliveryAuth) GenerateToken(c *fiber.Ctx) error {

	token, err := da.AuthUseCase.GetToken(c.IP())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
