package chamado

import (
	"destrella66/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var chamados []common.Chamado

	err := r.Db.Find(&chamados).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar chamado",
			"error":  err.Error(),
		})
	}

	return app.JSON(&fiber.Map{
		"Count": len(chamados),
		"Items": chamados,
	})
}
