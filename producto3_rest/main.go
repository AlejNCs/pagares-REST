package main

import (
	"log"
	"os"
    "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"producto3_rest/db"
	"producto3_rest/routes"
//"roducto3_rest/utils"
)

func main() {
	
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	db.Conectar()

	app := fiber.New()
    app.Use(cors.New())
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
    app.Static("/", "./public")
	app.Listen(":3000")

	log.Fatal(app.Listen(":" + port))
}
