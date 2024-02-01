package admins

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var admins []common.Admin

	err := r.Db.Find(&admins).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar admins",
			"error":  err.Error(),
		})
	}

	return app.JSON(&fiber.Map{
		"Count": len(admins),
		"items": admins,
	})
}
