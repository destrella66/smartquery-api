package admins

import (
	"destrella66/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var admin common.Admin

	err := app.BodyParser(&admin)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "admin esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Create(&admin).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar admin",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(&fiber.Map{
		"Items": admin,
	})
}
