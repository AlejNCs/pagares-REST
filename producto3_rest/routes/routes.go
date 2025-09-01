package routes

import (
	"github.com/gofiber/fiber/v2"
	"producto3_rest/controllers"
	//roducto3_rest/utils"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/pagares")

	api.Post("/", controllers.RegistrarPagare)
	api.Get("/:id", controllers.ConsultarPagare)
	api.Put("/:id", controllers.ModificarPagare)
	api.Delete("/:id", controllers.BorrarPagare)
app.Get("/api/pagares/pdf/:id", controllers.DescargarPDF)



}

