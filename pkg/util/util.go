package util

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2/utils"
	"io"
	"mime/multipart"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// UuidV4 生成一个 UUID
func UuidV4() string {
	return utils.UUIDv4()
}

// ReadFileWithTimeout 读取文件内容并设置超时
func ReadFileWithTimeout(file *multipart.FileHeader, duration time.Duration) ([]byte, error) {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// 打开文件
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.New("error closing file: " + closeErr.Error()) // 返回关闭文件时的错误
		}
	}()

	// 创建一个缓冲区来读取文件
	data := make([]byte, file.Size)

	// 使用通道来处理读取操作
	done := make(chan struct{})
	go func() {
		_, err = io.ReadFull(f, data)
		if err != nil {
			// 在读取操作中返回错误
			done <- struct{}{}
			return
		}
		done <- struct{}{}
	}()

	// 等待读取完成或超时
	select {
	case <-ctx.Done():
		return nil, ctx.Err() // 超时返回错误
	case <-done:
		return data, nil // 返回读取的文件内容
	}
}

func TileIsValid(x, y int32, zoom int8) bool {
	if zoom > 32 || zoom < 0 {
		return false
	}
	worldTileSize := int32(1) << uint32(zoom)
	if x < 0 || x >= worldTileSize ||
		y < 0 || y >= worldTileSize {
		return false
	}
	return true
}

// GetFileExtension 处理本地文件路径和 HTTP URL，根据 needDot 参数控制是否带点返回文件后缀
func GetFileExtension(filePath string, needDot bool) string {
	var ext string

	// 尝试解析是否是 URL
	parsedUrl, err := url.Parse(filePath)
	if err == nil && parsedUrl.Scheme != "" {
		// 是一个 URL，使用 path.Ext 提取后缀
		ext = path.Ext(parsedUrl.Path)
	} else {
		// 否则，认为是本地文件路径，使用 filepath.Ext
		ext = filepath.Ext(filePath)
	}

	if !needDot {
		// 如果不需要点，去掉前面的点
		ext = strings.TrimPrefix(ext, ".")
	}

	return ext
}

// GetFileName 处理本地文件路径和 HTTP URL，提取文件名，并根据 needExt 控制是否带后缀
func GetFileName(filePath string, needExt bool) string {
	var fileName string

	// 尝试解析是否是 URL
	parsedUrl, err := url.Parse(filePath)
	if err == nil && parsedUrl.Scheme != "" {
		// 是一个 URL，使用 path.Base 提取文件名
		fileName = path.Base(parsedUrl.Path)
	} else {
		// 否则，认为是本地文件路径，使用 filepath.Base
		fileName = filepath.Base(filePath)
	}

	if !needExt {
		// 如果不需要后缀，去掉扩展名
		ext := filepath.Ext(fileName)                // 提取扩展名
		fileName = strings.TrimSuffix(fileName, ext) // 去掉后缀
	}

	return fileName
}
