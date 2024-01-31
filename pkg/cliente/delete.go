package cliente

import (
	"anliben/teste/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Delete(app *fiber.Ctx) error {
	var cliente common.Cliente

	id := app.Params("id")

	err := r.Db.Delete(&cliente, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao deletar cliente",
			"error":  err.Error(),
		})
	}

	return app.JSON(cliente)
}
