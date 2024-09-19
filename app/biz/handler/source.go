package handler

import (
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util/response"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetItems godoc
// @Summary Get items from a source
// @Description Retrieves items from a specified source type
// @Tags source
// @Accept json
// @Produce json
// @Param sourceType path string true "Source type"
// @Param key query string false "Key for filtering items"
// @Success 200 {object} response.SuccessResponse{data=[]source.Item}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceType}/item [get]
func GetItems(ctx *fiber.Ctx) error {
	sourceType, err := validSourceType(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", sourceType))
	}
	key := ctx.Query("key", "")
	
	resp, err := sourceClient.GetItem(context.Background(), &source.GetItemRequest{
		SourceType: sourceType,
		Path:       "",
		Key:        key,
	})
	if err != nil {
		return err
	}

	return response.Success(ctx, resp.Items)
}

// AddItem godoc
// @Summary Add an item to the source
// @Description Adds a new item to the specified source
// @Tags source
// @Accept json
// @Produce json
// @Param sourceType path string true "Source type"
// @Param item body source.Item true "Item to add"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceType}/add [post]
func AddItem(ctx *fiber.Ctx) error {
	// TODO: Implement item addition logic
	return response.Success(ctx, "")
}

// PresignedUpload godoc
// @Summary Get pre-signed upload URL
// @Description Generates a pre-signed URL for uploading an item
// @Tags source
// @Accept json
// @Produce json
// @Param sourceType path string true "Source type"
// @Param path query string false "Path for the item"
// @Param name query string true "Name of the item"
// @Success 200 {object} response.SuccessResponse{data=map[string]string{uploadURL=string}}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceType}/presigned-upload [post]
func PresignedUpload(ctx *fiber.Ctx) error {
	sourceType, err := validSourceType(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", sourceType))
	}
	path := ctx.Query("path", "")
	name := ctx.Query("name", "")
	resp, err := sourceClient.PresignedUpload(context.Background(), &source.PresignedUploadResquest{
		SourceType: sourceType,
		Path:       path,
		Name:       name,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	return response.Success(ctx, fiber.Map{
		"uploadURL": resp.Url,
	})
}

// Publish godoc
// @Summary Publish an item
// @Description Publishes an item to the specified source
// @Tags source
// @Accept json
// @Produce json
// @Param sourceType path string true "Source type"
// @Param item body source.Item true "Item to publish"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceType}/publish [post]
func Publish(ctx *fiber.Ctx) error {
	// TODO: Implement database insertion logic
	// TODO: Implement record addition logic
	return nil
}

// validSourceType is an internal function and doesn't need Swagger documentation
func validSourceType(ctx *fiber.Ctx) (source.SourceType, error) {
	sourceType, err := source.SourceTypeFromString(ctx.Params("sourceType"))
	if err != nil {
		return sourceType, err
	}
	return sourceType, nil
}