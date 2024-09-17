package handler

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util/response"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetItems(ctx *fiber.Ctx) error {
	sourceType, err := validSourceType(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupport", sourceType))
	}
	key := ctx.Query("key", "")
	items, err := sourceClient.GetItem(context.Background(), &source.GetItemRequest{
		SourceType: sourceType,
		Path:       "",
		Key:        key,
	})
	if err != nil {
		return err
	}

	return response.Success(ctx, items.Items)
}

func AddItem(ctx *fiber.Ctx) error {

	return response.Success(ctx, "")
}

func PresignedUpload(ctx *fiber.Ctx) error {
	sourceType, err := validSourceType(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupport", sourceType))
	}
	path := ctx.Query("path", "")
	name := ctx.Query("name", "")
	upload, err := sourceClient.PresignedUpload(context.Background(), &source.PresignedUploadResquest{
		SourceType: sourceType,
		Path:       path,
		Name:       name,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	return response.Success(ctx, fiber.Map{
		"uploadURL": upload.Url,
	})
}

func Publish(ctx *fiber.Ctx) error {
	// 先入库

	// 然后添加记录

	return nil
}

func validSourceType(ctx *fiber.Ctx) (source.SourceType, error) {
	sourceType, err := source.SourceTypeFromString(ctx.Params("sourceType"))
	if err != nil {
		return sourceType, err
	}
	return sourceType, nil
}
