package db

import (
	"realworld-go-fiber/adapter/repository/sql/model"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	config util.Config
	logger port.Logger
	db     *gorm.DB
}

func New(config util.Config, logger port.Logger) (*DB, error) {
	var err error
	database := &DB{
		config: config,
		logger: logger,
	}
	if database.db, err = database.connect(); err != nil {
		return nil, err
	}
	if err := database.autoMigrate(); err != nil {
		return nil, err
	}
	return database, nil
}

func (db *DB) DB() *gorm.DB {
	return db.db
}

func (db *DB) connect() (*gorm.DB, error) {
	gdb, err := gorm.Open(postgres.Open(db.config.PostgresSource), &gorm.Config{
		// Logger:         logger.Default.LogMode(logger.Silent),
		DryRun:         false,
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	pg, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	err = pg.Ping()
	if err != nil {
		return nil, err
	}
	return gdb, nil
}

func (db *DB) autoMigrate() error {
	return db.db.AutoMigrate(&model.User{}, &model.Article{}, &model.Comment{}, &model.Tag{})
}
