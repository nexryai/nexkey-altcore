package activitypub

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
)

func ProcessFollowActivity(activity FollowActivity) error {
	logger.Info(fmt.Sprintf("Followed %s => %s", activity.Actor, activity.Object))
	return nil
}

func ProcessCreateActivity(activity CreateActivity) error {
	logger.Info(fmt.Sprintf("Note %s < \"%s\"", activity.Actor, activity.Object.MisskeyContent))
	return nil
}
