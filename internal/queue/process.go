package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"lab.sda1.net/nexryai/altcore/internal/activitypub"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	deliverProcessor "lab.sda1.net/nexryai/altcore/internal/queue/processor/deliver"
	inboxProcessor "lab.sda1.net/nexryai/altcore/internal/queue/processor/inbox"
	"net/http"
	"time"
)

const (
	TypeDeliverJob             = "ap:deliver"
	TypeFollowActivityInboxJob = "ap:TypeFollowActivityInbox"
)

type DeliverJob struct {
	TargetInbox    string
	UserId         string
	CreateActivity activitypub.CreateActivity
}

type ProcessFollowActivityJob struct {
	Activity activitypub.FollowActivity
	Request  http.Request
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
	deliverService := deliverProcessor.ActivityPubDeliverService{
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

func NewFollowActivityInboxTask(job ProcessFollowActivityJob) (*asynq.Task, error) {
	payload, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}
	// task options can be passed to NewTask, which can be overridden at enqueue time.
	return asynq.NewTask(TypeFollowActivityInboxJob, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute)), nil
}

func HandleFollowActivityInboxQueue(ctx context.Context, t *asynq.Task) error {
	var p ProcessFollowActivityJob
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	err := inboxProcessor.ProcessFollowActivity(p.Activity, &p.Request)
	if err != nil {
		logger.ErrorWithDetail("Failed to process", err)
		return err
	}

	return nil
}
