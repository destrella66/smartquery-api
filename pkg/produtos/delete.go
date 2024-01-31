package produtos

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Delete(app *fiber.Ctx) error {
	var produto common.Produto

	id := app.Params("id")

	err := r.Db.Delete(&produto, id).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao deletar produto",
			"error":  err.Error(),
		})
	}

	return app.JSON(produto)
}
