package chamado

import (
	"anliben/teste/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetOne(app *fiber.Ctx) error {
	var chamado common.Chamado

	id := app.Params("id")

	err := r.Db.First(&chamado, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar chamado",
			"error":  err.Error(),
		})
	}

	return app.JSON(chamado)
}
