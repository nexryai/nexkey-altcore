package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/system"
	"xorm.io/xorm"
)

func GetEngine() (*xorm.Engine, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DB.User,
		config.DB.Pass,
		config.DB.Host,
		config.DB.Port,
		config.DB.DB)

	engine, err := xorm.NewEngine("postgres", dbUrl)
	if err != nil {
		panic(system.UnableToCreateDatabaseSession)
	}

	engine.ShowSQL(true)

	return engine, nil
}

func GetGormEngine() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DB.Host,
		config.DB.User,
		config.DB.Pass,
		config.DB.DB,
		config.DB.Port)

	// Create a new Gorm DB instance
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
