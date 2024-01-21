package boot

import (
	"fmt"
	"github.com/hibiken/asynq"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/queue"
)

func QueueProcessDaemon() {
	redisAddr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: config.QueueConcurrency,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.TypeDeliverJob, queue.HandleDeliverQueue)
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		logger.FatalWithDetail("could not run asynq server: %v", err)
	}
}
