package router

import (
	"cloud_tinamic/app/biz/handler"
	"cloud_tinamic/app/biz/middleware"
	"github.com/gofiber/fiber/v2"
)

func registerSource(api fiber.Router) {
	source := api.Group("/source/:sourceCategory")

	source.Get("/items", middleware.Protected(), handler.GetItems)
	source.Get("/presigned-upload", middleware.Protected(), middleware.AuthRoutePermission(), handler.PresignedUpload)
	source.Get("/upload", middleware.Protected(), handler.Upload)
	source.Post("/publish", middleware.Protected(), middleware.AuthRoutePermission(), handler.Publish)
	source.Post("/add", middleware.Protected(), handler.AddItem)

	source.Post("/new-folder", middleware.Protected(), handler.NewFolder)
}
