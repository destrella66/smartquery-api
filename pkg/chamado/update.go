package chamado

import (
	"anliben/teste/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var chamado common.Chamado
	var foo common.Chamado
	id := app.Params("id")

	err := app.BodyParser(&foo)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "chamado esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Model(&chamado).
		Where("id = ?", id).
		UpdateColumns(common.Chamado{
			Titulo:    foo.Titulo,
			Descricao: foo.Descricao,
		}).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao fazer update no chamado",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(foo)
}
