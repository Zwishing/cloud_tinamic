package response

import "github.com/gofiber/fiber/v2"

// Response returns a JSON response with the given status, code, data, and message.
// It sets the HTTP status and returns the JSON response in one call.
func Response(ctx *fiber.Ctx, httpStatus int, code int, data interface{}, msg string) error {
	return ctx.Status(httpStatus).JSON(fiber.Map{"code": code, "data": data, "msg": msg})
}

// Success returns a successful response with status 200 OK.
// It's a convenience wrapper around Response for common success scenarios.
func Success(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, 200, data, "返回成功")
}

// Fail returns a failure response with status 404 Not Found.
// It's a convenience wrapper around Response for common failure scenarios.
func Fail(ctx *fiber.Ctx, msg string) error {
	return Response(ctx, fiber.StatusNotFound, 404, fiber.Map{}, msg)
}
