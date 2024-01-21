package store

import (
	"lab.sda1.net/nexryai/altcore/internal/core"
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
)

type FollowStore struct {
	Type   enum.FollowType
	UserId string
}

func (store *FollowStore) Create(followerId string) error {
	engine, err := db.GetEngine()
	if err != nil {
		return err
	}

	follow := entities.Following{
		Id:         core.GenId(),
		FollowerId: followerId,
		FolloweeId: store.UserId,
	}

	sql := engine.Table("following")
	_, err = sql.Insert(follow)
	if err != nil {
		return err
	}

	return nil
}

func (store *FollowStore) FindAll(joinUserInfo bool) ([]string, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return nil, err
	}

	var results []entities.Following

	sql := engine.Table("following")
	if store.Type == enum.Followees {
		sql.Where("\"followerId\" = ?", store.UserId)
	} else if store.Type == enum.Followers {
		sql.Where("\"followeeId\" = ?", store.UserId)
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	if joinUserInfo {
		// TODO: join user info
	}

	err = sql.Find(&results)
	if err != nil {
		return nil, err
	}

	var resultArr []string

	for _, result := range results {
		if store.Type == enum.Followees {
			resultArr = append(resultArr, result.FolloweeId)
		} else if store.Type == enum.Followers {
			resultArr = append(resultArr, result.FollowerId)
		}
	}

	return resultArr, nil

}
