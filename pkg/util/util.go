package util

import (
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"mime/multipart"
)

func UuidV4() string {
	return utils.UUIDv4()
}

// ReadFile 读取文件内容
func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	return io.ReadAll(f)
}
