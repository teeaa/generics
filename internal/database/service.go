package database

import (
	"fmt"

	"github.com/teeaa/generics/internal/config"
	"github.com/teeaa/generics/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func New(conf *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", conf.Database.Host, conf.Database.User, conf.Database.Password, conf.Database.Name, conf.Database.Port)
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db := &Database{db: gormDb}
	err = db.migrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (d *Database) migrate() error {
	return d.db.AutoMigrate(
		&models.Address{},
		&models.Dob{},
		&models.Name{},
	)
}
