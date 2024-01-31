package setup

import (
	"destrella66/pkg/admins"
	"destrella66/pkg/chamado"
	"destrella66/pkg/cliente"
	"destrella66/pkg/common"
	"destrella66/pkg/menu"
	"destrella66/pkg/produtos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup() {

	app := fiber.New()

	db, _ := common.OpenConnection()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "",
		AllowCredentials: true,
	}))

	chamado.RegisterRoutes(app, db)
	cliente.RegisterRoutes(app, db)
	produtos.RegisterRoutes(app, db)
	admins.RegisterRoutes(app, db)
	menu.RegisterRoutes(app, db)

	app.Listen(":3000")
}
