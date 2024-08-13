package route

import (
	"notification-service/internal/config"
	httpdelivery "notification-service/internal/delivery/http"
	middleware "notification-service/internal/delivery/http/midlleware"
	"notification-service/internal/repository/postgresql"
	"notification-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	// Initialize http
	app.Use(middleware.ErrorHandlerMiddleware)

	userRepo := postgresql.NewUserRepository(cfg.DB)
	userService := service.NewUserService(userRepo, cfg)
	userController := httpdelivery.NewUserController(userService)

	// Public routes
	app.Post("/users", userController.Register)
	app.Get("/users/:id", userController.GetUserByID)

}
