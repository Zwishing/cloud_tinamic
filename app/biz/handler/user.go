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

// Login
// @Summary 用户登录
// @Description 用户使用账号和密码登录
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param   requestBody  body   model.LoginRequest  true  "登录请求参数"
// @Success 200 {object} map[string]interface{} "返回用户ID和JWT Token"
// @Failure 400 {object} response.FailResponse "请求参数错误或登录失败"
// @Router /login [post]
func Login(ctx *fiber.Ctx) error {
	//获取参数
	var signin model.LoginRequest

	if err := ctx.BodyParser(&signin); err != nil {
		return response.Fail(ctx, err.Error())
	}

	if err := validate.ValidateRequestBody(signin); err != nil {
		return response.Fail(ctx, err.Error())
	}

	category, err := user.UserCategoryFromString(strings.ToUpper(signin.Category))
	if err != nil {
		return response.Fail(ctx, err.Error())
	}

	// 验证登录
	login, err := userClinet.Login(context.Background(), &user.LoginRequest{
		Username:     signin.UserAccount,
		Password:     signin.Password,
		UserCategory: category,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}

	// 生成JWT Token
	token, err := jwt.ReleaseToken(login.UserId)
	if err != nil {
		return response.Fail(ctx, err.Error())
	}

	return response.Success(ctx, fiber.Map{
		"userId": login.UserId,
		"token":  token,
	})
}

// Register 用户注册接口
// @Summary 注册一个新用户
// @Description 使用用户名和密码注册一个新用户，只有管理员有权限注册，设置角色，默认是游客角色
// @ID register-user
// @Accept  json
// @Produce  json
// @Param   loginAccount    object  true  "User Registration"
// @Success 200 {object} map[string]string
// @Router /v1/user/register [post]
func Register(ctx *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Fail(ctx, err.Error())
	}

	if err := validate.ValidateRequestBody(req); err != nil {
		return response.Fail(ctx, err.Error())
	}

	category, err := user.UserCategoryFromString(strings.ToUpper(req.Category))
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	_, err = userClinet.Register(context.Background(), &user.RegisterRequest{
		Username:     req.UserAccount,
		Password:     req.Password,
		UserCategory: category,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	return response.Success(ctx, "注册成功")
}

func CurrentUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(string)
	profile, err := userClinet.Info(context.Background(), &user.InfoRequest{
		UserId: userId,
	})
	if err != nil {
		return response.Fail(ctx, err.Error())
	}
	return response.Success(ctx, profile)
}
