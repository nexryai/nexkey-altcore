package services

import (
	"fmt"
	"github.com/hibiken/asynq"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/queue"
)

func AddCreateActivityToDeliverQueue(job queue.DeliverJob) error {
	task, err := queue.NewDeliverTask(job)
	if err != nil {
		logger.ErrorWithDetail("could not create task: %v", err)
		return err
	}

	const redisAddr = "127.0.0.1:6379"
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	info, err := client.Enqueue(task)
	if err != nil {
		logger.ErrorWithDetail("could not enqueue task: %v", err)
	}

	logger.Info(fmt.Sprintf("enqueued task: id=%s queue=%s", info.ID, info.Queue))
	return nil
}
