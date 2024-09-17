package router

import "github.com/gofiber/fiber/v2"

func RegisterAPI(api fiber.Router) {
	registerUser(api)
	registerSource(api)
}
