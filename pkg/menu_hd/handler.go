// api/menu/handler.go
package menuDB

import (
	"destrella66/pkg/common"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Menu common.Menu

var menuHD = []Menu{
	{
		Id:    1,
		Label: "Home",
		Link:  "dashboard",
		Icon:  "po-icon po-icon-chart-area",
	},
	{
		Id:    2,
		Label: "Chamados",
		Link:  "chamados",
		Icon:  "po-icon po-icon-news",
	},
	
	{
		Id:    4,
		Label: "Produtos",
		Link:  "produtos",
		Icon:  "po-icon po-icon-pallet-full",
	},
	{
		Id:    5,
		Label: "Voltar",
		Link:  "sair",
		Icon:  "po-icon po-icon-exit",
	},
}

func (r handler) GetMenuHD(app *fiber.Ctx) error {

	response := map[string]interface{}{
		"items": menuHD,
	}
	menuJSON, err := json.Marshal(response)

	if err != nil {
		return app.Status(fiber.StatusInternalServerError).SendString("Houve um erro no endpoint do menu do HelpDesl")
	}
	return app.Status(fiber.StatusOK).Send(menuJSON)
}
