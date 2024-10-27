package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/teeaa/generics/internal"
	"github.com/teeaa/generics/internal/database"
	"github.com/teeaa/generics/internal/models"
)

type Op[T models.DataType] struct{ s *Service }

func NewOps[T models.DataType](s *Service) Op[T] {
	return Op[T]{s: s}
}

func (t Op[T]) Get(ctx context.Context, item models.Item[T]) (*T, error) {
	res, err := database.NewOps[T](t.s.db).Get(ctx, item)
	if err != nil {
		slog.Error("Get", slog.Any("err", err))
		if errors.Is(err, internal.ErrNotFound) {
			return nil, internal.ErrNotFound
		}

		return nil, internal.ErrServerError
	}
	return res, nil
}
