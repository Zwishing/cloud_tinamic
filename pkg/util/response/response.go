package response

import "github.com/gofiber/fiber/v2"

// Response returns a JSON response with the given status, code, data, and message.
// It sets the HTTP status and returns the JSON response in one call.
func Response(ctx *fiber.Ctx, httpStatus int, data interface{}, msg string) error {
	return ctx.Status(httpStatus).JSON(fiber.Map{"code": httpStatus, "data": data, "msg": msg})
}

func SuccessWithOK(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data, "返回成功")
}

func SuccessWithCreated(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data, "请求成功")
}
func SuccessWithNoContent(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"code": fiber.StatusNoContent, "msg": "处理成功，无返回数据"})
}

// Fail returns a failure response with status 404 Not Found.
// It's a convenience wrapper around Response for common failure scenarios.
func Fail(ctx *fiber.Ctx, httpStatus int, msg string) error {
	return ctx.Status(httpStatus).JSON(fiber.Map{"code": httpStatus, "msg": msg})
}

func FailWithBadRequest(ctx *fiber.Ctx, msg string) error {
	return Fail(ctx, fiber.StatusBadRequest, msg)
}

func FailWithNotFound(ctx *fiber.Ctx, msg string) error {
	return Fail(ctx, fiber.StatusNotFound, msg)
}

func FailWithNonAuthoritativeInformation(ctx *fiber.Ctx, msg string) error {
	return Fail(ctx, fiber.StatusNonAuthoritativeInformation, msg)
}

func FailWithExpectation(ctx *fiber.Ctx, msg string) error {
	return Fail(ctx, fiber.StatusExpectationFailed, msg)
}

func FailWithInternalServerError(ctx *fiber.Ctx, msg string)error{
	return Fail(ctx, fiber.StatusInternalServerError, msg)
}
