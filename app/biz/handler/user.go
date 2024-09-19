package handler

import (
	"cloud_tinamic/app/biz/model"
	"cloud_tinamic/kitex_gen/base/user"
	"cloud_tinamic/pkg/util/jwt"
	"cloud_tinamic/pkg/util/response"
	"cloud_tinamic/pkg/util/validate"
	"context"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// Login godoc
// @Summary User login
// @Description Authenticate a user with account and password
// @Tags user
// @Accept json
// @Produce json
// @Param requestBody body model.LoginRequest true "Login credentials"
// @Success 200 {object} response.SuccessResponse{data=map[string]interface{userId=string,token=string}} "Returns user ID and JWT Token"
// @Failure 400 {object} response.ErrorResponse "Invalid request parameters or login failure"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /v1/user/login [post]
func Login(ctx *fiber.Ctx) error {
	var signin model.LoginRequest

	// Parse and validate request body
	if err := ctx.BodyParser(&signin); err != nil {
		return response.Fail(ctx, "Invalid request body")
	}

	if err := validate.ValidateRequestBody(signin); err != nil {
		return response.Fail(ctx, err.Error())
	}

	// Convert and validate user category
	category, err := user.UserCategoryFromString(strings.ToUpper(signin.Category))
	if err != nil {
		return response.Fail(ctx, "Invalid user category")
	}

	// Authenticate user
	resp, err := userClient.Login(context.Background(), &user.LoginRequest{
		Username:     signin.UserAccount,
		Password:     signin.Password,
		UserCategory: category,
	})
	if err != nil {
		return response.Fail(ctx, "Authentication failed")
	}

	// Generate JWT token
	token, err := jwt.ReleaseToken(resp.UserId)
	if err != nil {
		return response.Fail(ctx, "Failed to generate token")
	}

	// Return success response with user ID and token
	return response.Success(ctx, fiber.Map{
		"userId": resp.UserId,
		"token":  token,
	})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password. Only administrators can register and set roles. Default role is guest.
// @Tags user
// @Accept json
// @Produce json
// @Param requestBody body model.RegisterRequest true "User registration details"
// @Success 200 {object} response.SuccessResponse{data=string} "Registration successful"
// @Failure 400 {object} response.ErrorResponse "Invalid request parameters"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /v1/user/register [post]
func Register(ctx *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Fail(ctx, "Invalid request body")
	}

	if err := validate.ValidateRequestBody(req); err != nil {
		return response.Fail(ctx, err.Error())
	}

	category, err := user.UserCategoryFromString(strings.ToUpper(req.Category))
	if err != nil {
		return response.Fail(ctx, "Invalid user category")
	}

	_, err = userClient.Register(context.Background(), &user.RegisterRequest{
		Username:     req.UserAccount,
		Password:     req.Password,
		UserCategory: category,
	})
	if err != nil {
		return response.Fail(ctx, "Registration failed")
	}

	return response.Success(ctx, "Registration successful")
}

// CurrentUser godoc
// @Summary Get current user information
// @Description Retrieve information about the currently authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.SuccessResponse{data=user.User} "User information"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /v1/user/current [get]
func CurrentUser(ctx *fiber.Ctx) error {
	userId := ctx.Query("userId", "")
	if userId == ""{
		userId = ctx.Locals("userId").(string)
	}
	
	resp, err := userClient.Info(context.Background(), &user.InfoRequest{
		UserId: userId,
	})
	if err != nil {
		return response.Fail(ctx, "Failed to fetch user information")
	}
	return response.Success(ctx, resp.User)
}
