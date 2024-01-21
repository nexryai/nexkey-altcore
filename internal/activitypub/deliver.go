package activitypub

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/security"
	"lab.sda1.net/nexryai/altcore/internal/core/utils"
	"lab.sda1.net/nexryai/altcore/internal/services/xaccount"
	"net/http"
	"time"
)

type ActivityPubDeliverService struct {
	TargetInbox string
	UserId      string
}

func (params *ActivityPubDeliverService) DeliverCreateActivity(activity *CreateActivity) error {
	// SafeURLじゃないなら拒否
	if !security.IsSafeUrl(params.TargetInbox) {
		logger.Warn("discard this job (reason: !IsSafeUrl)")
		return nil
	}

	activity.Context = activityContext
	activity.ActivityType = "Create"

	keyringService := xaccount.KeyringService{
		UserId: params.UserId,
	}

	privateKey, err := keyringService.GetPrivateKeyPem()
	if err != nil {
		return err
	}

	apRequestService := &ActivityPubRequestService{
		Url:     params.TargetInbox,
		Headers: []Header{{Name: "Host", Value: utils.GetHostFromUrl(params.TargetInbox)}},
		Body:    activity,
		Method:  "POST",
	}

	httpRequest := apRequestService.ToHttpRequest()

	// 署名のための情報
	signService := SignatureService{
		PrivateKeyPem: privateKey,
		KeyId:         fmt.Sprintf("%s/users/%s", config.URL, params.UserId),
		Request:       httpRequest,
	}

	// httpRequestに署名する
	err = signService.Sign()
	if err != nil {
		return err
	}

	fmt.Println("Signed Headers:")
	for name, header := range httpRequest.Header {
		fmt.Printf("%s: %s\n", name, header)
	}

	// 実行
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	response, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusGone {
		logger.Info("discard this job (reason: 410 Gone)")
		return nil
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to deliver (target server fault: status %s)", response.Status)
	}

	logger.Info(fmt.Sprintf("Note Delivered: %s", params.TargetInbox))
	return nil
}
