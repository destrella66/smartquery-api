package produtos

import (
	"destrella66/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Update(app *fiber.Ctx) error {
	var produto common.Produto
	var foo common.Produto
	id := app.Params("id")

	err := app.BodyParser(&foo)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "produto esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Model(&produto).
		Where("id = ?", id).
		UpdateColumns(common.Produto{
			Nome:  produto.Nome,
			Preco: produto.Preco,
		}).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao fazer update no produto",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(foo)
}
