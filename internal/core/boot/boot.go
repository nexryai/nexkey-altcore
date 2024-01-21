package boot

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"os"
)

func Init() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "localhost"
	}

	fmt.Printf("Starting nexkey worker on %s...\n", hostName)
	fmt.Printf("\nPID: %d\nUID: %d\n\n", os.Getpid(), os.Getuid())

	// 未知の連合先からもリクエストを受け付けまくるものなのにrootとして動かすのは危険すぎるので拒否
	if os.Getuid() == 0 {
		if os.Getenv("UNSAFE_MODE") == "I_UNDERSTAND_WHAT_I_AM_DOING" {
			logger.Warn(">>>>>> Running the server as root is very dangerous and must not do! <<<<<<")
			logger.Warn("DO NOT USE IN PRODUCTION")
		} else {
			logger.Error("Running the server as root is very dangerous and must not do!")
			os.Exit(1)
		}
	}

	logger.ProgressInfo("Loading config...")
	if len(config.Secret) < 32 {
		logger.Error("config.Secret must be at least 32 characters")
		os.Exit(1)
	}

	logger.ProgressOk()

	logger.ProgressInfo("Checking DB...")
	// ToDO: エラーハンドリング
	instance.GetInstanceMeta()

	logger.ProgressOk()

}
