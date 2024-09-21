package util

import (
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"mime/multipart"
)

func UuidV4() string {
	return utils.UUIDv4()
}

// ReadFile reads the content of a file
func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
