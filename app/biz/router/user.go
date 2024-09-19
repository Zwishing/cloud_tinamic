package router

import (
	"cloud_tinamic/app/biz/handler"
	middleware "cloud_tinamic/app/biz/middleware"
	"github.com/gofiber/fiber/v2"
)

func registerUser(api fiber.Router) {
	user := api.Group("/user")

	//user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
	user.Get("/currentUser", middleware.Protected(), handler.CurrentUser)
}
