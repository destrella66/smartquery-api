package cliente

import (
	"anliben/teste/pkg/common"

	"github.com/gofiber/fiber/v2"
)

func (r handler) GetAll(app *fiber.Ctx) error {
	var clientes []common.Cliente

	err := r.Db.Find(&clientes).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar clientes",
			"error":  err.Error(),
		})
	}

	return app.JSON(&fiber.Map{
		"Count": len(clientes),
		"Items": clientes,
	})
}
