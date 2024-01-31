package admins

import (
	"destrella66/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var admin common.Admin
	var foo common.Admin
	id := app.Params("id")

	err := app.BodyParser(&foo)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "admin esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Model(&admin).
		Where("id = ?", id).
		UpdateColumns(common.Admin{
			Nome:           foo.Nome,
			Endereco:       foo.Endereco,
			Genero:         foo.Genero,
			NumeroEndereco: foo.NumeroEndereco,
			Telefone:       foo.Telefone,
			Estado:         foo.Estado,
		}).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao fazer update no admin",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(foo)
}
