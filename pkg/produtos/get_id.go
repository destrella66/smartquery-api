package produtos

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetOne(app *fiber.Ctx) error {
	var produto common.Produto

	id := app.Params("id")

	err := r.Db.First(&produto, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar produto",
			"error":  err.Error(),
		})
	}

	return app.JSON(produto)
}
