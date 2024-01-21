package instance

import (
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
	"lab.sda1.net/nexryai/altcore/internal/db/kv"
)

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
