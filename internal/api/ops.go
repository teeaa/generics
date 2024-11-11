package api

import (
	"errors"
	"log/slog"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/teeaa/generics/internal"
	"github.com/teeaa/generics/internal/models"
	"github.com/teeaa/generics/internal/service"
)

type Op[T models.DataType, Req models.RequestType[T], Res models.ResponseType[T]] struct {
	// inject service.Service to be able to call it
	s *service.Service
}

// Constructor
func NewOps[T models.DataType, Req models.RequestType[T], Res models.ResponseType[T]](s *service.Service) Op[T, Req, Res] {
	return Op[T, Req, Res]{s: s}
}

func (t Op[T, Req, Res]) Create(c *fuego.ContextWithBody[Req]) (*Res, error) {
	req, err := c.Body() // req is of type Req
	if err != nil {
		slog.Error("Create", slog.Any("err", err))
	}

	baseType := req.ToModel() // baseType is of type T

	// Create an interface to be able to cast it to Item[T]
	var tt any
	tt = baseType
	item := tt.(models.Item[T])

	slog.Info("Create", slog.Any("item", item))

	id := uuid.New().String()
	txID := c.PathParam("txid")

	base := &models.Base{
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		ID:            &id,
		TransactionID: &txID,
	}

	tt = item.SetBase(base)
	item = tt.(models.Item[T])

	slog.Info("Create", slog.Any("item", item))

	ctx := c.Context()

	res, err := service.NewOps[T](t.s).Create(ctx, item)
	if err != nil {
		return nil, handleError(err)
	}

	var r interface{}
	r = res
	resItem := r.(models.ResponseItem[T, Res])

	return resItem.ToResponse(), nil
}

func (t Op[T, Req, Res]) Get(c *fuego.ContextWithBody[Req]) (*Res, error) {
	slog.Info("Get")
	ctx := c.Context()
	req, err := c.Body()
	if err != nil {
		slog.Error("Get", slog.Any("err", err))
	}

	var item any
	item = req.ToModel()

	res, err := service.NewOps[T](t.s).Get(ctx, item.(models.Item[T]))
	if err != nil {
		return nil, handleError(err)
	}

	var r interface{}
	r = res
	resItem := r.(models.ResponseItem[T, Res])

	return resItem.ToResponse(), nil
}

func handleError(err error) error {
	switch {
	case errors.Is(err, internal.ErrNotFound):
		return fuego.HTTPError{Err: err, Status: 404, Detail: "Not Found"}
	default:
		return fuego.HTTPError{Err: err, Status: 500, Detail: "Internal Server Error"}
	}
}
