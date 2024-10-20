package main

import (
	"log"
	"os"

	"github.com/B6137151/GDZ-Commerce/internal/domain/repository"
	"github.com/B6137151/GDZ-Commerce/internal/dto/http/controller"
	"github.com/B6137151/GDZ-Commerce/internal/dto/http/middleware"
	"github.com/B6137151/GDZ-Commerce/internal/infrastructure/database"
	"github.com/B6137151/GDZ-Commerce/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	// Initialize the PostgreSQL database
	db, err := database.NewPostgresDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Middleware for logging and error recovery
	app.Use(logger.New())  // Log requests
	app.Use(recover.New()) // Recover from panics

	// Initialize Store repository, service, and controller
	storeRepo := repository.NewStoreRepository(db)
	storeService := service.NewStoreService(storeRepo)
	storeController := controller.NewStoreController(storeService)

	// Initialize Auth repository, service, and controller for Admin/User authentication
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	// Authentication routes for Admin and User
	app.Post("/admin/login", authController.AdminLogin)
	app.Post("/user/login", authController.UserLogin)
	app.Post("/admin/register", authController.RegisterAdmin) // Add admin registration
	app.Post("/user/register", authController.RegisterUser)   // Add user registration

	// Refresh token route
	app.Post("/auth/refresh", authController.RefreshToken) // New route for refreshing tokens

	// Store Management - Only Admin Access
	storeGroup := app.Group("/stores")
	storeGroup.Use(middleware.JWTProtected())       // Apply JWT protection to all store routes
	storeGroup.Use(middleware.RequireRole("admin")) // Restrict to admins only

	// Store management endpoints (admin-only access)
	storeGroup.Post("/", storeController.CreateStore)
	storeGroup.Get("/:id", storeController.GetStore)
	storeGroup.Put("/:id", storeController.UpdateStore)
	storeGroup.Delete("/:id", storeController.DeleteStore)

	// Secure routes for Admin and User (using JWT)
	adminGroup := app.Group("/admin")
	adminGroup.Use(middleware.JWTProtected())       // Protect admin routes with JWT
	adminGroup.Use(middleware.RequireRole("admin")) // Ensure only admins can access these routes
	adminGroup.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Admin Dashboard!")
	})

	userGroup := app.Group("/user")
	userGroup.Use(middleware.JWTProtected())      // Protect user routes with JWT
	userGroup.Use(middleware.RequireRole("user")) // Ensure only users can access these routes
	userGroup.Get("/profile", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to your User Profile!")
	})

	// Determine the port to use (default to 3000 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the server on the specified port
	log.Printf("Server is starting on port %s...", port)
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
