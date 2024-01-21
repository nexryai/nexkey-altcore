package auth

import (
	"github.com/judwhite/argon2"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/services/baselib"
)

func PasswordIsOk(userId string, checkPassword string) bool {
	userService := baselib.UserService{
		LocalOnly: true,
	}

	i, err := userService.GetProfile(userId)
	if err != nil {
		logger.ErrorWithDetail("Error getting user profile", err)
		panic(err)
	}

	err = argon2.CompareHashAndPassword(i.PasswordHash, []byte(checkPassword))
	if err == nil {
		return true
	}

	return false
}
