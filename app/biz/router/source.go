package router

import (
	"cloud_tinamic/app/biz/handler"
	"cloud_tinamic/app/biz/middleware"
	"github.com/gofiber/fiber/v2"
)

func registerSource(api fiber.Router) {
	source := api.Group("/source/:sourceCategory")
	source.Get("/homeItems", middleware.Protected(), handler.GetHomeItems)
	source.Get("/nextItems", middleware.Protected(), handler.GetNextItems)
	source.Get("/previousItems", middleware.Protected(), handler.GetPreviousItems)

	source.Get("/presignedUpload", middleware.Protected(), middleware.AuthRoutePermission(), handler.PresignedUpload)
	source.Post("/upload", middleware.Protected(), handler.Upload)
	source.Post("/publish", handler.Publish)
	source.Post("/add", middleware.Protected(), handler.AddItem)

	source.Post("/newFolder", middleware.Protected(), handler.NewFolder)
	source.Delete("deleteItems", middleware.Protected(), handler.DeleteItems)

}
