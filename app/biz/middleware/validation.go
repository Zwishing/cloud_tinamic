package middleware

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ValidSourceCateogry() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category, err := source.SourceCategoryFromString(strings.ToUpper(ctx.Params("sourceCategory")))
		if err != nil {
			return response.FailWithBadRequest(ctx, "sourceCategory is unsupported")
		}
		ctx.Locals("sourceCategory", category)

		return ctx.Next()
	}
}

