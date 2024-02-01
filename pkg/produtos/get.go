package produtos

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var produtos []common.Produto

	err := r.Db.Find(&produtos).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar produtos",
			"error":  err.Error(),
		})
	}

	return app.JSON(&fiber.Map{
		"Count": len(produtos),
		"items": produtos,
	})
}
