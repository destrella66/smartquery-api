package cliente

import (
	"anliben/teste/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var cliente common.Cliente

	err := app.BodyParser(&cliente)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "cliente esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Create(&cliente).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar cliente",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(&fiber.Map{
		"Items": cliente,
	})
}
