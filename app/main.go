package main

import (
	"cloud_tinamic/app/biz/router"
	"cloud_tinamic/app/logger"
	conf "cloud_tinamic/config"
	"fmt"
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
// @BasePath /api/v1
func main() {
	app := InitApp()
	app.Use(cors.New())

	// 使用自定义日志库
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger.Log,
	}))

	api := app.Group("/v1")
	router.RegisterAPI(api)

	//设置端口监听
	if err := app.Listen(fmt.Sprintf(":%d", conf.GetConfigInstance().GetInt("server.port"))); err != nil {
		logger.Log.Fatal().Err(err).Msg("Fiber app error")
	}

}

type App struct {
	*fiber.App
	Hasher hashing.Driver
	//Session *session.Session
}

func InitApp() *App {
	app := &App{
		App: fiber.New(),
		//Hasher: hashing.New(config.Conf.GetHasherConfig()),
		//Session: session.New(CONFIGFILE.GetSessionConfig()),
	}
	return app
}
