package xdrive

import (
	"database/sql"
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/db/kv"
)

type DriveService struct {
	FileId    string
	LocalOnly bool
}

func (param *DriveService) FindOne() (entities.DriveFile, error) {
	var result entities.DriveFile
	cacheKey := "drive/file/" + param.FileId

	cacheExist := kv.GetKvCache(cacheKey, &result)

	if !cacheExist {
		database, err := db.GetGormEngine()
		if err != nil {
			return result, err
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

		sql := database.Table("drive_file")
		sql.Where("\"id\" = ?", param.FileId)
		if param.LocalOnly {
			sql.Where("\"userHost\" is NULL")
		}

		err = sql.First(&result).Error
		if err != nil {
			return result, err
		}

		kv.StoreKvCache(cacheKey, result)
	}

	return result, nil
}

func (param *DriveService) FindAllAndMap(fileIds []string) (*map[string]entities.DriveFile, error) {
	if len(fileIds) == 0 || param.FileId != "" {
		panic(system.InvalidParamsOnServiceCall)
	}

	var result map[string]entities.DriveFile
	cacheKey := "drive/file/" + fmt.Sprintf("%v", fileIds)

	cacheExists := kv.GetKvCache(cacheKey, &result)
	if !cacheExists {
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

		var files []entities.DriveFile

		sql := database.Table("drive_file")
		sql.Where("id IN (?)", fileIds)
		if param.LocalOnly {
			sql.Where("host is NULL")
		}

		if err := sql.Find(&files).Error; err != nil {
			return nil, err
		}

		result = make(map[string]entities.DriveFile)
		for _, file := range files {
			result[file.Id] = file
		}

		kv.StoreKvCache(cacheKey, result)
	}

	return &result, nil
}
