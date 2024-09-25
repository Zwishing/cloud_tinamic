package main

import (
	"cloud_tinamic/app/biz/router"
	"cloud_tinamic/app/logger"
	conf "cloud_tinamic/config"
	"fmt"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	hashing "github.com/thomasvvugt/fiber-hashing"
)

// @title Tinamic服务API
// @version 1.0
// @description
// @termsOfService

// @contact.name API Support
// @contact.url
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /v1
func main() {
	app := InitApp()

	// Apply middleware
	app.Use(cors.New())
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))

	// Set up API routes
	api := app.Group("/v1")
	router.RegisterAPI(api)

	// Start server
	port := conf.GetConfigInstance().GetInt("server.port")
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to run Cloud-Tinamic App")
	}
}

type App struct {
	*fiber.App
	Hasher hashing.Driver
}

func InitApp() *App {
	return &App{
		App: fiber.New(fiber.Config{
			ReadTimeout:  3 * time.Minute, // 设置超时时间
			WriteTimeout: 3 * time.Minute,
			// 可以设置大文件上传的最大 body size，比如 100 MB
			BodyLimit: 300 * 1024 * 1024, // 100MB
		}),
		// Initialize Hasher if needed
		// Hasher: hashing.New(config.Conf.GetHasherConfig()),
	}
}
