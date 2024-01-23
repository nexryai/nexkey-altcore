package migration

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lopezator/migrator"
	"github.com/nexryai/visualog"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
)

const (
	migrationTable = "altcore_migrations"
)

var versionMigrations = []migrator.Option{
	migrator.TableName(migrationTable),
	migrationsToAltcore,
}

var (
	log   = logger.GetLogger("migration")
	dbUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.DB)
)

func runVersionMigrations(db *sql.DB, logger *visualog.Logger) error {
	m, err := migrator.New(versionMigrations...)
	if err != nil {
		return err
	}

	err = m.Migrate(db)
	if err != nil {
		return err
	}

	log.Info("Migrations completed!")
	return nil
}

func RunDatabaseMigration() error {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.FatalWithDetail("Failed to connect to database", err)
		return err
	}

	err = runVersionMigrations(db, log)
	if err != nil {
		log.FatalWithDetail("Failed to run version migrations", err)
		return err
	}

	return nil
}

func InitDB() error {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.FatalWithDetail("Failed to connect to database", err)
		return err
	}

	m, err := migrator.New(migrator.TableName(migrationTable), migrationsToInitMisskey)
	if err != nil {
		log.FatalWithDetail("Failed to create migrator", err)
		return err
	}

	err = m.Migrate(db)
	if err != nil {
		log.Fatal("Failed... >_<")
		log.FatalWithDetail("Failed to init DB", err)
		return err
	}

	log.Info("Initiation completed!")
	return nil
}
