package api

import (
	"errors"
	"log/slog"

	"github.com/go-fuego/fuego"
	"github.com/teeaa/generics/internal"
	"github.com/teeaa/generics/internal/models"
	"github.com/teeaa/generics/internal/service"
)

type Op[T models.DataType] struct {
	// inject service.Service to be able to call it
	s *service.Service
}

// Constructor
func NewOps[T models.DataType](s *service.Service) Op[T] {
	return Op[T]{s: s}
}

func (t Op[T]) Create(c *fuego.ContextWithBody[T]) (*T, error) {
	req, err := c.Body() // req is of type T
	if err != nil {
		slog.Error("Create", slog.Any("err", err))
	}

	// Create an interface to be able to cast it to Item[T]
	var tt any
	tt = req
	item := tt.(models.Item[T])

	slog.Info("Create", slog.Any("item", item))

	// id := uuid.New().String()
	// txID := c.PathParam("txid")

	// base := &models.Base{
	// 	CreatedAt:     time.Now(),
	// 	UpdatedAt:     time.Now(),
	// 	ID:            &id,
	// 	TransactionID: &txID,
	// }

	// tt = item.SetBase(base)
	// item = tt.(models.Item[T])

	// slog.Info("Create", slog.Any("item", item))

	// ctx := c.Context()

	// res, err := service.NewOps[T](t.s).Create(ctx, item)
	// if err != nil {
	// 	return nil, handleError(err)
	// }

	// return res, nil
	return nil, nil
}

func (t Op[T]) Get(c *fuego.ContextWithBody[T]) (*T, error) {
	slog.Info("Get")
	ctx := c.Context()
	req, err := c.Body()
	if err != nil {
		slog.Error("Get", slog.Any("err", err))
	}

	var item any
	item = req

	res, err := service.NewOps[T](t.s).Get(ctx, item.(models.Item[T]))
	if err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func handleError(err error) error {
	switch {
	case errors.Is(err, internal.ErrNotFound):
		return fuego.HTTPError{Err: err, Status: 404, Detail: "Not Found"}
	default:
		return fuego.HTTPError{Err: err, Status: 500, Detail: "Internal Server Error"}
	}
}
