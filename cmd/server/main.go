package main

import (
	"log"

	"github.com/B6137151/GDZ-Commerce/internal/domain/repository"
	"github.com/B6137151/GDZ-Commerce/internal/dto/http/controller"
	"github.com/B6137151/GDZ-Commerce/internal/infrastructure/database"
	"github.com/B6137151/GDZ-Commerce/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.NewPostgresDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	app := fiber.New()

	storeRepo := repository.NewStoreRepository(db)
	storeService := service.NewStoreService(storeRepo)
	storeController := controller.NewStoreController(storeService)

	app.Post("/stores", storeController.CreateStore)
	app.Get("/stores/:id", storeController.GetStore)
	app.Put("/stores/:id", storeController.UpdateStore)
	app.Delete("/stores/:id", storeController.DeleteStore)

	log.Println("Server is starting...")
	log.Fatal(app.Listen(":3000"))
}
