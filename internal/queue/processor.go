package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"net/http"
	"time"
)

const (
	TypeDeliverJob = "ap:deliver"
	TypeInboxJob   = "ap:inbox"
)

type DeliverJob struct {
	TargetInbox    string
	UserId         string
	CreateActivity activitypub.CreateActivity
}

type InboxJobPayload struct {
	UserID  string
	Request http.Request
}

func NewDeliverTask(job DeliverJob) (*asynq.Task, error) {
	payload, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeDeliverJob, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func HandleDeliverQueue(ctx context.Context, t *asynq.Task) error {
	var p DeliverJob
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	deliverService := activitypub.ActivityPubDeliverService{
		TargetInbox: p.TargetInbox,
		UserId:      p.UserId,
	}

	err := deliverService.DeliverCreateActivity(&p.CreateActivity)
	if err != nil {
		logger.ErrorWithDetail("Failed to deliver", err)
		return err
	}

	return nil
}
