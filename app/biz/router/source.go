package router

import (
	"cloud_tinamic/app/biz/handler"
	"cloud_tinamic/app/biz/middleware"
	"github.com/gofiber/fiber/v2"
)

func registerSource(api fiber.Router) {
	source := api.Group("/source/:sourceType")

	source.Get("/items", middleware.Protected(), middleware.AuthRoutePermission(), handler.GetItems)
	source.Get("/presigned-upload", middleware.Protected(), middleware.AuthRoutePermission(), handler.PresignedUpload)
	source.Post("/publish", middleware.Protected(), middleware.AuthRoutePermission(), handler.Publish)
	source.Post("/add",middleware.Protected(),middleware.AuthRoutePermission(),handler.AddItem)
}
