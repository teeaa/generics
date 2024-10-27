package database

import (
	"context"
	"errors"

	"github.com/teeaa/generics/internal"
	"github.com/teeaa/generics/internal/models"
	"gorm.io/gorm"
)

type Op[T models.DataType] struct {
	*Database
}

func NewOps[T models.DataType](db *Database) Op[T] {
	return Op[T]{db}
}

func (s Op[T]) Get(ctx context.Context, item models.Item[T]) (*T, error) {
	var ret T
	q := s.db.WithContext(ctx).Where(item).Order("created_at DESC")

	res := q.Debug().First(&ret)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrNotFound
		}
		return nil, internal.ErrNotFound
	}

	return &ret, nil
}
