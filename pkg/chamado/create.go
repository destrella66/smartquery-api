package chamado

import (
	"anliben/teste/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var chamado common.Chamado

	err := app.BodyParser(&chamado)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "chamado esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Create(&chamado).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar chamado",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(&fiber.Map{
		"Items": chamado,
	})
}
