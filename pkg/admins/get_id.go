package admins

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetOne(app *fiber.Ctx) error {
	var admin common.Admin

	id := app.Params("id")

	err := r.Db.First(&admin, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar admin",
			"error":  err.Error(),
		})
	}

	return app.JSON(admin)
}
