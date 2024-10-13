package handler

import (
	"cloud_tinamic/app/biz/model"
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/kitex_gen/service/collection"
	"cloud_tinamic/pkg/util/response"
	"cloud_tinamic/pkg/util/validate"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetCollections(ctx *fiber.Ctx) error {
	return response.SuccessWithOK(ctx, "")
}

func Publish(ctx *fiber.Ctx) error {
	sourceCategory, ok := ctx.Locals("sourceCategory").(source.SourceCategory)
	if !ok {
		return response.FailWithInternalServerError(ctx, "Invalid source category")
	}

	var publishRequest model.PublishRequest
	if err := ctx.BodyParser(&publishRequest); err != nil {
		return response.FailWithBadRequest(ctx, "Invalid request body")
	}

	if err := validate.ValidateRequestBody(publishRequest); err != nil {
		return response.FailWithBadRequest(ctx, err.Error())
	}

	serviceCategory, err := collection.ServiceCategoryFromString(strings.ToUpper(publishRequest.ServiceCategory))
	if err != nil {
		return response.FailWithBadRequest(ctx, "Invalid service category")
	}

	err = collectionClient.Publish(ctx.Context(), &collection.PublishRequest{
		SourceCategory:  sourceCategory,
		SourceKey:       publishRequest.Key,
		ServiceCategory: serviceCategory,
	})

	if err != nil {
		return response.FailWithInternalServerError(ctx, "Failed to publish: "+err.Error())
	}
	return response.SuccessWithOK(ctx, "Publication successful")
}
