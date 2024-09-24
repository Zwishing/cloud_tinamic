package handler

import (
	"cloud_tinamic/app/biz/model"
	"cloud_tinamic/kitex_gen/base"
	"cloud_tinamic/kitex_gen/data/source"
	"cloud_tinamic/kitex_gen/data/storage"
	"cloud_tinamic/pkg/util"
	"cloud_tinamic/pkg/util/response"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetHomeItems(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}

	resp, err := sourceClient.GetHomeItems(ctx.Context(), &source.GetHomeItemsRequest{
		SourceCategory: category,
	})
	if err != nil {
		return err
	}
	return response.Success(ctx, fiber.Map{
		"key":   resp.Key,
		"items": model.Items(resp.Items),
	})

}

// GetNextItems godoc
// @Summary Get items from a source
// @Description Retrieves items from a specified source type
// @Tags source
// @Accept json
// @Produce json
// @Param sourceCategory path string true "Source type"
// @Param key query string false "Key for filtering items"
// @Success 200 {object} response.SuccessResponse{data=[]source.Item}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/items [get]
func GetNextItems(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}
	key := ctx.Query("key", "")

	resp, err := sourceClient.GetNextItems(ctx.Context(), &source.GetItemsRequest{
		SourceCategory: category,
		Key:            key,
	})
	if err != nil {
		return err
	}

	return response.Success(ctx, fiber.Map{
		"key":   resp.Key,
		"items": model.Items(resp.Items),
	})
}

func GetPreviousItems(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}
	key := ctx.Query("key", "")

	resp, err := sourceClient.GetPreviousItems(ctx.Context(), &source.GetItemsRequest{
		SourceCategory: category,
		Key:            key,
	})
	if err != nil {
		return err
	}
	return response.Success(ctx, fiber.Map{
		"key":   resp.Key,
		"items": model.Items(resp.Items),
	})

}

// AddItem godoc
// @Summary Add an item to the source
// @Description Adds a new item to the specified source
// @Tags source
// @Accept json
// @Produce json
// @Param sourceCategory path string true "Source type"
// @Param item body source.Item true "Item to add"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/add [post]
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
// @Param sourceCategory path string true "Source type"
// @Param path query string false "Path for the item"
// @Param name query string true "Name of the item"
// @Success 200 {object} response.SuccessResponse{data=map[string]string{uploadURL=string}}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/presignedUpload [post]
func PresignedUpload(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}
	path := ctx.Query("path", "")
	name := ctx.Query("name", "")
	resp, err := sourceClient.PresignedUpload(ctx.Context(), &source.PresignedUploadResquest{
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

// Upload godoc
// @Summary Upload a file to the source
// @Description Uploads a file to the specified source type
// @Tags source
// @Accept multipart/form-data
// @Produce json
// @Param sourceCategory path string true "Source Category"
// @Param file formData file true "File to upload"
// @Param parentId formData string false "Parent ID for the file"
// @Param name formData string true "Name of the file"
// @Success 200 {object} response.SuccessResponse{data=map[string]string{sourceId=string}}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/upload [post]
func Upload(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}
	var uploadReq model.UploadRequest
	if err = ctx.BodyParser(&uploadReq); err != nil {
		return response.Fail(ctx, "请求体无效")
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		return response.Fail(ctx, "无法从表单获取文件")
	}
	readFile, err := util.ReadFileWithTimeout(file, 1*time.Minute)
	if err != nil {
		return response.Fail(ctx, "读取文件失败")
	}
	resp, err := sourceClient.Upload(ctx.Context(), &source.UploadRequest{
		SourceCategory: category,
		Key:            uploadReq.Key,
		Name:           uploadReq.Name,
		Size:           file.Size,
		FileData:       readFile,
	})
	if err != nil {
		return response.Fail(ctx, "文件上传失败")
	}
	return response.Success(ctx, fiber.Map{
		"key": resp.Key,
	})
}

// NewFolder godoc
// @Summary Create a new folder
// @Description Creates a new folder in the specified source
// @Tags source
// @Accept json
// @Produce json
// @Param sourceCategory path string true "Source type"
// @Param newFolder body model.NewFolderRequest true "New folder details"
// @Success 200 {object} response.SuccessResponse{data=model.Item}
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/newFolder [post]
func NewFolder(ctx *fiber.Ctx) error {
	category, err := validSourceCategory(ctx)
	if err != nil {
		return response.Fail(ctx, fmt.Sprintf("sourceCategory %s is unsupported", category))
	}

	var newFolderReq model.NewFolderRequest
	if err := ctx.BodyParser(&newFolderReq); err != nil {
		return response.Fail(ctx, "Invalid request body")
	}

	resp, err := sourceClient.CreateFolder(ctx.Context(), &source.CreateFolderRequest{
		SourceCategory: category,
		Key:            newFolderReq.Key,
		Name:           newFolderReq.Name,
		Path:           newFolderReq.Path,
	})
	if err != nil || resp.Base.Code != base.Code_SUCCESS {
		return response.Fail(ctx, fmt.Sprintf("Failed to create folder,%s", resp.Base.Msg))
	}

	return response.Success(ctx, model.Item(resp.Item))
}

// Publish godoc
// @Summary Publish an item
// @Description Publishes an item to the specified source
// @Tags source
// @Accept json
// @Produce json
// @Param sourceCategory path string true "Source type"
// @Param item body source.Item true "Item to publish"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 503 {object} response.ErrorResponse
// @Router /v1/source/{sourceCategory}/publish [post]
func Publish(ctx *fiber.Ctx) error {
	// TODO: Implement database insertion logic
	// TODO: Implement record addition logic
	resp, err := geoClient.VectorStorage(ctx.Context(), &storage.StoreRequest{
		Schema: "public",
		Table:  "sandy",
		Name:   "aa",
		Url:    "http://39.101.164.253:9000/vector/石漠化监测数据.zip",
		Ext:    "zip",
	})
	if err != nil {
		return err
	}
	if resp.Base.Code == base.Code_SUCCESS {
		return response.Success(ctx, resp.Base.Msg)
	} else {
		return response.Fail(ctx, resp.Base.Msg)
	}
}

func DeleteItems(ctx *fiber.Ctx) error {
	keys := strings.Split(ctx.Query("key", ""), ",")
	resp, err := sourceClient.DeleteItems(ctx.Context(), &source.DeleteItemsRequest{Keys: keys})
	if err != nil {
		return err
	}
	if resp.Base.Code != base.Code_SUCCESS {
		return response.Fail(ctx, resp.Base.Msg)
	}
	return response.Success(ctx, resp.Base.Msg)

}

// validSourceCategory is an internal function and doesn't need Swagger documentation
func validSourceCategory(ctx *fiber.Ctx) (source.SourceCategory, error) {
	category, err := source.SourceCategoryFromString(strings.ToUpper(ctx.Params("sourceCategory")))
	if err != nil {
		return category, err
	}
	return category, nil
}
