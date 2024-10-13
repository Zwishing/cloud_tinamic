package router

import (
	"cloud_tinamic/app/biz/handler"
	"cloud_tinamic/app/biz/middleware"
	"github.com/gofiber/fiber/v2"
)

func registerVectorService(api fiber.Router) {
	service := api.Group("/service/:sourceCategory", middleware.Protected(), middleware.ValidSourceCateogry())

	service.Post("/publish", handler.Publish)

	service.Get("/collections", handler.GetCollections)
	service.Get("/collections/:collectionId")
	service.Get("/collections/:collectionId/items")
	service.Get("/collections/:collectionId/items/:itemId")

	service.Get("/collections/:collectionId/tiles")
	service.Get("/collections/:collectionId/tiles/:tileMatrixSetId")
	service.Get("/collections/:collectionId/tiles/:tileMatrixSetId/:x/:y/:z")
	service.Get("/collections/:collectionId/tiles/:tileMatrixSetId/tilejson.json")
}
