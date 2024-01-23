package v12api

import "github.com/gofiber/fiber/v2"

type ApiErrorResp struct {
	Error string `json:"error"`
	Id    string `json:"id"`
}

type ApiError struct {
	Resp       *ApiErrorResp
	StatusCode int
}

var (
	ApiErrPasswordInvalid = &ApiError{
		Resp: &ApiErrorResp{
			Error: "Invalid password",
			Id:    "932c904e-9460-45b7-9ce6-7ed33be7eb2c",
		},
		StatusCode: 403,
	}
	ApiErrUserNotFound = &ApiError{
		Resp: &ApiErrorResp{
			Error: "User not found",
			Id:    "6cc579cc-885d-43d8-95c2-b8c7fc963280",
		},
		StatusCode: 400,
	}
)

func ReturnApiError(ctx *fiber.Ctx, err *ApiError) error {
	return ctx.Status(err.StatusCode).JSON(err.Resp)
}
