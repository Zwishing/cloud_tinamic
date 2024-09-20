package handler

import (
	"cloud_tinamic/app/biz/model"
	"cloud_tinamic/kitex_gen/base"
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/pkg/util/response"
	"context"
	"fmt"
	"strings"

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
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", category))
	}
	key := ctx.Query("sourceId", "")

	resp, err := sourceClient.GetItem(context.Background(), &source.GetItemRequest{
		SourceCategory: category,
		Path:           "",
		Key:            key,
	})
	if err != nil {
		return err
	}

	return response.Success(ctx, model.Items(resp.Items))
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
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", category))
	}
	path := ctx.Query("path", "")
	name := ctx.Query("name", "")
	resp, err := sourceClient.PresignedUpload(context.Background(), &source.PresignedUploadResquest{
		SourceCategory: category,
		Path:           path,
		Name:           name,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	return response.Success(ctx, fiber.Map{
		"uploadURL": resp.Url,
	})
}

func Upload(ctx *fiber.Ctx) error {
	// 接收文件
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", category))
	}
	file, err := ctx.FormFile("file")
	var uploadReq model.UploadRequest
	if err := ctx.BodyParser(&uploadReq); err != nil {
		return response.Fail(ctx, "Invalid request body")
	}
	readFile, err := util.ReadFile(file)
	if err != nil {
		// TODO
		return err
	}
	_, err = sourceClient.Upload(context.Background(), &source.UploadRequest{
		SourceCategory: category,
		ParentId:       uploadReq.ParentId,
		Name:           uploadReq.Name,
		Size:           file.Size,
		FileData:       readFile,
	})
	if err != nil {
		return err
	}
	return response.Success(ctx, "")
}

func NewFolder(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceType %s is unsupported", category))
	}

	var newFolderReq model.NewFolderRequest
	if err := ctx.BodyParser(&newFolderReq); err != nil {
		return response.Fail(ctx, "Invalid request body")
	}

	resp, err := sourceClient.CreateFolder(context.Background(), &source.CreateFolderRequest{
		SourceCategory: category,
		ParentId:       newFolderReq.ParentId,
		Name:           newFolderReq.Name,
		Path:           newFolderReq.Path,
	})
	if err != nil {
		return response.Fail(ctx, "Failed to create folder")
	}
	if resp.Base.Code != base.Code_SUCCESS {
		return response.Fail(ctx, "Failed to create folder")
	}
	return response.Success(ctx, model.Item(resp.Item))
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

// validSourceCategory is an internal function and doesn't need Swagger documentation
func validSourceCategory(ctx *fiber.Ctx) (source.SourceCategory, error) {
	category, err := source.SourceCategoryFromString(strings.ToUpper(ctx.Params("sourceCategory")))
	if err != nil {
		return category, err
	}
	return category, nil
}
