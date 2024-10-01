package handler

import (
	"cloud_tinamic/pkg/util/response"
	"github.com/gofiber/fiber/v2"
)

func GetCollections(ctx *fiber.Ctx) error {
	return response.Success(ctx, "")
}
