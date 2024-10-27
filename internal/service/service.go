package service

import (
	"github.com/teeaa/generics/internal/config"
	"github.com/teeaa/generics/internal/database"
)

type Service struct {
	db *database.Database
}

func New(conf *config.Config, db *database.Database) *Service {
	return &Service{db: db}
}
