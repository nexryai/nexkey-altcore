package instance

import (
	"database/sql"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/db/kv"
)

func ShouldInitDB() bool {
	database, err := db.GetGormEngine()
	if err != nil {
		panic(err)
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

	// metaテーブルが存在するか
	if !database.Migrator().HasTable("meta") {
		return true
	}

	return false
}

func GetInstanceMeta() *entities.Meta {
	var meta entities.Meta

	cacheExist := kv.GetKvCache("meta", &meta)

	if !cacheExist {
		database, err := db.GetGormEngine()
		if err != nil {
			panic(err)
		}

		sql := database.Table("meta")
		sql.Where("id = ?", "x")
		err = sql.First(&meta).Error
		if err != nil {
			logger.ErrorWithDetail("Failed to get instance metadata", err)
			panic(err)
		}

		kv.StoreKvCache("meta", meta)
	}

	return &meta
}
