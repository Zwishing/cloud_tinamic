package middleware

import (
	"cloud_tinamic/app/logger"
	"cloud_tinamic/kitex_gen/base/auth"
	"cloud_tinamic/kitex_gen/base/auth/authservice"
	"cloud_tinamic/pkg/util/jwt"
	"cloud_tinamic/pkg/util/response"
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/gofiber/fiber/v2"
)

var (
	authClient authservice.Client
	once       sync.Once
)

// initAuthClient initializes the Kitex client for authentication
func initAuthClient() {
	once.Do(func() {
		var err error
		authClient, err = authservice.NewClient("base.auth.authservice", client.WithHostPorts("0.0.0.0:8811"))
		if err != nil {
			logger.Log.Fatal().Msgf("Failed to create auth client: %v", err)
		}
	})
}

// Protected middleware for JWT authentication
func Protected() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return response.Fail(ctx, "Missing Authorization header")
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return response.Fail(ctx, "Invalid Authorization header format")
		}

		claims, err := jwt.ValidateToken(tokenParts[1])
		if err != nil {
			return response.Fail(ctx, "Invalid or expired token")
		}

		ctx.Locals("userId", claims["userId"])
		return ctx.Next()
	}
}

// JWTErrorHandler handles JWT-related errors
func JWTErrorHandler(c *fiber.Ctx, err error) error {
	status := fiber.StatusUnauthorized
	message := "Invalid or expired JWT"

	if err.Error() == "Missing or malformed JWT" {
		status = fiber.StatusBadRequest
		message = "Missing or malformed JWT"
	}

	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}

// AuthRoutePermission middleware for route-based authorization
func AuthRoutePermission() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		initAuthClient() // Ensure auth client is initialized

		sub, ok := ctx.Locals("userId").(string)
		if !ok {
			return response.Fail(ctx, "User ID not found in context")
		}

		resp, err := authClient.Auth(context.Background(), &auth.AuthResquest{
			Sub: sub,
			Obj: ctx.Path(),
			Act: ctx.Method(),
		})
		if err != nil {
			return response.FailWithNonAuthoritativeInformation(ctx, fmt.Sprintf("Authorization error: %v", err))
		}
		if !resp.Allow {
			return response.FailWithNonAuthoritativeInformation(ctx, "Access denied")
		}

		return ctx.Next()
	}
}
