package setup

import (
	"anliben/teste/pkg/chamado"
	"anliben/teste/pkg/cliente"
	"anliben/teste/pkg/common"

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

	app.Listen(":3000")
}
