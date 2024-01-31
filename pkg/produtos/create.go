package produtos

import (
	"destrella66/pkg/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (r handler) Create(app *fiber.Ctx) error {
	var produto common.Produto

	err := app.BodyParser(&produto)

	if err != nil {
		return app.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"detail": "produto esta com algum parametro errado",
			"error":  err.Error(),
		})
	}

	err = r.Db.Create(&produto).Error

	if err != nil {
		return app.JSON(&fiber.Map{
			"detail": "erro ao buscar produto",
			"error":  err.Error(),
		})
	}

	return app.Status(http.StatusCreated).JSON(&fiber.Map{
		"Items": produto,
	})
}
