package inbox

import (
	"fmt"
	"github.com/nexryai/visualog"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"net/http"
)

func getInboxLogger() *visualog.Logger {
	return logger.GetLogger("inbox")
}

func ProcessFollowActivity(activity activitypub.FollowActivity, request *http.Request) error {
	log := getInboxLogger()

	/*
		signService := activitypub.SignatureService{
			PrivateKeyPem: "", //ToDo
			Request:       request,
		}
	*/

	log.Info(fmt.Sprintf("Followed %s => %s", activity.Actor, activity.Object))
	return nil
}

func ProcessCreateActivity(activity activitypub.CreateActivity, request *http.Request) error {
	log := getInboxLogger()
	log.Info(fmt.Sprintf("Note %s < \"%s\"", activity.Actor, activity.Object.MisskeyContent))
	return nil
}
