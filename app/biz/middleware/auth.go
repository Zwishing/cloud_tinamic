package middleware

import (
	"cloud_tinamic/app/logger"
	"cloud_tinamic/kitex_gen/base/auth"
	"cloud_tinamic/kitex_gen/base/auth/authservice"
	"cloud_tinamic/pkg/util/jwt"
	"cloud_tinamic/pkg/util/response"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gofiber/fiber/v2"
)

var (
	authClient authservice.Client
)

func init() {
	// 创建 Kitex 客户端
	var err error
	authClient, err = authservice.NewClient("base.auth.authservice", client.WithHostPorts("0.0.0.0:8811"))
	if err != nil {
		logger.Log.Fatal().Msgf("创建客户端失败: %v", err)
	}
}

func Protected() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// 从请求头中获取 Authorization 字段
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return response.Fail(ctx, "缺少 Authorization")
		}

		// 检查 Bearer token
		tokenString := authHeader[len("Bearer "):]
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			return response.Fail(ctx, "Authorization 验证失败")
		}

		ctx.Locals("userId", claims["userId"])

		return ctx.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

// AuthRoutePermission 验证路由权限
func AuthRoutePermission() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sub := ctx.Locals("userId").(string)
		obj := ctx.Path()
		act := ctx.Method()
		resp, err := authClient.Auth(context.Background(), &auth.AuthResquest{
			Sub: sub,
			Obj: obj,
			Act: act,
		})
		if err != nil {
			return err
		}
		if !resp.Allow {
			return fmt.Errorf("access denied")
		}

		return ctx.Next()
	}
}
