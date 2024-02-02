// api/menu/handler.go
package menu

import (
	"destrella66/pkg/common"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Menu common.Menu
type SubMenuItem common.SubMenuItem

var menu = []Menu{
	{
		Id:    1,
		Label: "Início",
		Link:  "home",
		Icon:  "po-icon-home",
	},
	{
		Id:    2,
		Label: "Perfil",
		Link:  "profile",
		Icon:  "po-icon-user",
	},
	{
		Id:    3,
		Label: "Início",
		Link:  "home",
		Icon:  "po-icon-home",
	},
	{
		Id:    4,
		Label: "Consulta Personalizada",
		Link:  "custom-query",
		Icon:  "po-icon-clipboard",
		SubItems: []common.SubMenuItem{
			{
				Label: "Portal do Admin",
				Link:  "admin",
				Icon:  "po-icon-user",
			},
			{
				Label: "Portal do Cliente",
				Link:  "cliente",
				Icon:  "po-icon-user",
			},
		},
	},
	{
		Id:    5,
		Label: "Outras Ações",
		Link:  "others",
		Icon:  "po-icon-more",
	},
	{
		Id:    6,
		Label: "HelpDesk",
		Link:  "helpdesk",
		Icon:  "po-icon-clipboard",
		SubItems: []common.SubMenuItem{
			{
				Label: "DashBoard",
				Link:  "helpdesk",
				Icon:  "po-icon-user",
			},
			{
				Label: "Produtos",
				Link:  "produtos",
				Icon:  "po-icon-user",
			},
			{
				Label: "Chamados",
				Link:  "chamados",
				Icon:  "po-icon-user",
			},
		},
	},
}

func (r handler) GetMenu(app *fiber.Ctx) error {

	response := map[string]interface{}{
		"items": menu,
	}
	menuJSON, err := json.Marshal(response)

	if err != nil {
		return app.Status(fiber.StatusInternalServerError).SendString("Houve um erro no endpoint do menu")
	}
	return app.Status(fiber.StatusOK).Send(menuJSON)
}
