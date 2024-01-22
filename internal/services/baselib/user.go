package baselib

import (
	"database/sql"
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
	database, err := db.GetGormEngine()
	if err != nil {
		return entities.User{}, err
	}

	var result entities.User
	sql := database.Preload("Avatar").Preload("Banner").Table("user")

	if userId != "" {
		sql.Where("id = ?", userId)
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	if param.LocalOnly {
		sql.Where("host is NULL")
	}

	err = sql.Find(&result).Error
	if err != nil {
		return entities.User{}, err
	}

	return result, nil
}

func (param *UserService) FindOneByName(userName string) (*entities.User, error) {
	database, err := db.GetGormEngine()
	if err != nil {
		return nil, err
	}

	dbInstance, err := database.DB()
	if err != nil {
		panic(err)
	}

	defer func(dbInstance *sql.DB) {
		err := dbInstance.Close()
		if err != nil {
			panic(err)
		}
	}(dbInstance)

	var result entities.User
	sql := database.Table("user")

	if userName != "" {
		sql.Where("\"usernameLower\" = ?", strings.ToLower(userName))
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	if param.LocalOnly {
		sql.Where("host is NULL")
	}

	err = sql.First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// 配列内にあるIDのユーザー情報をidをインデックスとするmapにして返す
func (param *UserService) FindAllAndMap(userIds []string) (*map[string]entities.User, error) {
	database, err := db.GetGormEngine()
	if err != nil {
		return nil, err
	}

	dbInstance, err := database.DB()
	if err != nil {
		panic(err)
	}

	defer func(dbInstance *sql.DB) {
		err := dbInstance.Close()
		if err != nil {
			panic(err)
		}
	}(dbInstance)

	var users []entities.User

	query := database.Table("user")
	query.Where("id IN (?)", userIds)
	if param.LocalOnly {
		query.Where("host is NULL")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	userInfoMap := make(map[string]entities.User)
	for _, user := range users {
		userInfoMap[user.Id] = user
	}

	return &userInfoMap, nil
}

func (param *UserService) GetProfile(userId string) (entities.UserProfile, error) {
	database, err := db.GetGormEngine()
	if err != nil {
		return entities.UserProfile{}, err
	}

	var result entities.UserProfile

	query := database.Table("user_profile")
	query.Where("\"userId\" = ?", userId)

	err = query.First(&result).Error
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
