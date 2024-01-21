package baselib

import (
	"lab.sda1.net/nexryai/altcore/internal/core/enum"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"lab.sda1.net/nexryai/altcore/internal/db"
	"lab.sda1.net/nexryai/altcore/internal/db/entities"
)

type FollowService struct {
	Type   enum.FollowType
	UserId string
}

func (params *FollowService) FindAll() ([]string, error) {
	engine, err := db.GetEngine()
	if err != nil {
		return nil, err
	}

	var results []entities.Following

	sql := engine.Table("following")
	if params.Type == enum.Followees {
		sql.Where("\"followerId\" = ?", params.UserId)
	} else if params.Type == enum.Followers {
		sql.Where("\"followeeId\" = ?", params.UserId)
	} else {
		panic(system.InvalidParamsOnServiceCall)
	}

	err = sql.Find(&results)
	if err != nil {
		return nil, err
	}

	var resultArr []string

	for _, result := range results {
		if params.Type == enum.Followees {
			resultArr = append(resultArr, result.FolloweeId)
		} else if params.Type == enum.Followers {
			resultArr = append(resultArr, result.FollowerId)
		}
	}

	return resultArr, nil

}

func (params *FollowService) CheckVisibility(noteVisibility enum.NoteVisibility, userId string) bool {
	followers, err := params.FindAll()
	if err != nil {
		panic(err)
	}

	switch noteVisibility {
	case enum.NoteVisibilityPublic:
		return true
	case enum.NoteVisibilityHome:
		return true
	case enum.NoteVisibilityFollowersOnly:
		for _, follower := range followers {
			if follower == userId {
				return true
			}
		}
		return false
	case enum.NoteVisibilityDirect:
		return false
	}

	return false
}
