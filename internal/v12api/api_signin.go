package v12api

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/services/auth"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
	"lab.sda1.net/nexryai/altcore/internal/v12api/schema"
	"time"
)

var (
	ErrorPasswordInvalid = errors.New("invalid password")
	ErrorUserNotFound    = errors.New("invalid username")
	ErrorInvalidParams   = errors.New("invalid params")
)

func SignIn(req *schema.SignInRequest) (schema.SignInResp, error) {
	logger.Debug(fmt.Sprintf("login attemp: %s", req.Username))

	if req.Username == "" || req.Password == "" {
		return schema.SignInResp{}, ErrorInvalidParams
	}

	userService := baselib.UserService{
		LocalOnly: true,
	}

	i, err := userService.FindOneByName(req.Username)
	if err != nil || i.Id == "" {
		return schema.SignInResp{}, ErrorUserNotFound
	}

	// Check password
	if auth.PasswordIsOk(i.Id, req.Password) {
		claims := jwt.MapClaims{
			"accountToken": i.Token,
			"id":           i.Id,
			"exp":          time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		t, err := token.SignedString([]byte(config.Secret))
		if err != nil {
			logger.ErrorWithDetail("Error signing JWT token", err)
			panic(err)
		}

		return schema.SignInResp{
			Token:  t,
			UserId: i.Id,
		}, nil

	}

	return schema.SignInResp{}, ErrorPasswordInvalid
}
