package baselib

import (
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/core/utils"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"strings"
)

type UserService struct {
	LocalOnly bool
}

func (param *UserService) FindOne(userId string) (entities.User, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return entities.User{}, err
	}

	var result entities.User
	sql := engine.Table("user")

	if userId != "" {
		sql.Where("id = ?", userId)
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	if param.LocalOnly {
		sql.Where("host is NULL")
	}

	_, err = sql.Get(&result)
	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}

func (param *UserService) FindOneByName(userName string) (*entities.User, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return nil, err
	}

	var result entities.User
	sql := engine.Table("user")

	if userName != "" {
		sql.Where("usernameLower = ?", strings.ToLower(userName))
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	if param.LocalOnly {
		sql.Where("host is NULL")
	}

	_, err = sql.Get(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// 配列内にあるIDのユーザー情報をidをインデックスとするmapにして返す
func (param *UserService) FindAllAndMap(userIds []string) (*map[string]entities.User, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return nil, err
	}

	var users []entities.User

	sql := engine.Table("user")
	sql.In("id", userIds)
	if param.LocalOnly {
		sql.Where("host is NULL")
	}

	if err := sql.Find(&users); err != nil {
		return nil, err
	}

	userInfoMap := make(map[string]entities.User)
	for _, user := range users {
		userInfoMap[user.Id] = user
	}

	return &userInfoMap, nil
}

func (param *UserService) GetProfile(userId string) (entities.UserProfile, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return entities.UserProfile{}, err
	}

	var result entities.UserProfile

	sql := engine.Table("user_profile")
	sql.Where("\"userId\" = ?", userId)

	_, err = sql.Get(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (param *UserService) GetSharedInboxes(userIds []string) ([]string, error) {
	var result []string

	userInfos, err := param.FindAllAndMap(userIds)
	if err != nil {
		return result, err
	}

	for _, u := range *userInfos {
		if !utils.ContainsString(result, u.SharedInbox) {
			result = append(result, u.SharedInbox)
		}
	}

	return result, nil
}
