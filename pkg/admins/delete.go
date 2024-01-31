package admins

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Delete(app *fiber.Ctx) error {
	var admin common.Admin

	id := app.Params("id")

	err := r.Db.Delete(&admin, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao deletar admin",
			"error":  err.Error(),
		})
	}

	return app.JSON(admin)
}
