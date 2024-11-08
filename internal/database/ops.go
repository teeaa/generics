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

func (s Op[T]) Create(ctx context.Context, item models.Item[T]) error {
	res := s.db.WithContext(ctx).Save(item)

	if res.Error != nil {
		errors.Join(res.Error, errors.New("failed to create item"))
	}

	return nil
}

func (s Op[T]) Get(ctx context.Context, item models.Item[T]) (*T, error) {
	var ret T
	q := s.db.WithContext(ctx).Where(item).Order("created_at DESC")

	res := q.Debug().First(&ret)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, internal.ErrNotFound
		}
		return nil, errors.Join(res.Error, errors.New("failed to fetch item"))
	}

	return &ret, nil
}
