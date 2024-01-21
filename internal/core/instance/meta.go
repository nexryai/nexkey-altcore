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
		engine, err := db.GetEngine()
		if err != nil {
			panic(err)
		}

		sql := engine.Table("meta")
		sql.Where("id = ?", "x")
		_, err = sql.Get(&meta)
		if err != nil {
			logger.ErrorWithDetail("Failed to get instance metadata", err)
			panic(err)
		}

		kv.StoreKvCache("meta", meta)
	}

	return &meta
}
