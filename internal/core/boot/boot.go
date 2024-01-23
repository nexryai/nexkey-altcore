package boot

import (
	"fmt"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/instance"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"lab.sda1.net/nexryai/altcore/internal/db/migration"
	"os"
)

func Init() {
	log := logger.GetLogger("boot")

	hostName, err := os.Hostname()
	if err != nil {
		hostName = "localhost"
	}

	fmt.Printf("Starting nexkey worker on %s...\n", hostName)
	fmt.Printf("\nPID: %d\nUID: %d\n\n", os.Getpid(), os.Getuid())

	// 未知の連合先からもリクエストを受け付けまくるものなのにrootとして動かすのは危険すぎるので拒否
	if os.Getuid() == 0 {
		if os.Getenv("UNSAFE_MODE") == "I_UNDERSTAND_WHAT_I_AM_DOING" {
			log.Warn(">>>>>> Running the server as root is very dangerous and must not do! <<<<<<")
			log.Warn("DO NOT USE IN PRODUCTION")
		} else {
			log.Error("Running the server as root is very dangerous and must not do!")
			os.Exit(1)
		}
	}

	log.ProgressInfo("Loading config...")
	if len(config.Secret) < 32 {
		logger.Error("config.Secret must be at least 32 characters")
		os.Exit(1)
	}

	log.ProgressOk()

	log.ProgressInfo("Checking DB...")
	if instance.ShouldInitDB() {
		log.ProgressInfo("Initializing database...")
		err = migration.InitDB()
		if err != nil {
			log.ErrorWithDetail("Failed to init database :(", err)
			os.Exit(1)
		}
		log.ProgressOk()
	}

	err = migration.RunDatabaseMigration()
	if err != nil {
		log.ErrorWithDetail("Failed to run database migration", err)
		os.Exit(1)
	}

	m := instance.GetInstanceMeta()
	logger.Info(fmt.Sprintf("DB is OK: %s", m.Name))

	log.ProgressOk()
}
