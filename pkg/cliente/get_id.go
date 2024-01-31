package cliente

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetOne(app *fiber.Ctx) error {
	var cliente common.Cliente

	id := app.Params("id")

	err := r.Db.First(&cliente, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar cliente",
			"error":  err.Error(),
		})
	}

	return app.JSON(cliente)
}
