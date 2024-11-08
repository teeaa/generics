package api

// import (
// 	"errors"
// 	"log/slog"
// 	"time"

// 	"github.com/go-fuego/fuego"
// 	"github.com/google/uuid"
// 	"github.com/teeaa/generics/internal"
// 	"github.com/teeaa/generics/internal/models"
// 	"github.com/teeaa/generics/internal/service"
// )

// type Op[T models.DataType] struct {
// 	s       *service.Service
// 	subtype string
// }

// // Constructor
// func NewOps[T models.DataType](s *service.Service, subtype string) Op[T] {
// 	return Op[T]{s: s, subtype: subtype}
// }

// func (t Op[T]) Create(c *fuego.ContextWithBody[T]) (*T, error) {
// 	ctx := c.Context()
// 	req, err := c.Body() // req is of type T
// 	if err != nil {
// 		slog.Error("Create", slog.Any("err", err))
// 	}

// 	var tt any
// 	tt = req
// 	item := tt.(models.Item[T])

// 	id := uuid.New().String()
// 	txID := c.PathParam("txid")

// 	base := &models.Base{
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 		ID:            &id,
// 		TransactionID: &txID,
// 	}

// 	tt = item.SetBase(base)
// 	item = tt.(models.Item[T])

// 	res, err := service.NewOps[T](t.s).Create(ctx, item)
// 	if err != nil {
// 		return nil, handleError(err)
// 	}
// 	return res, nil
// }

// func (t Op[T]) Get(c *fuego.ContextWithBody[T]) (*T, error) {
// 	slog.Info("Get")
// 	ctx := c.Context()
// 	req, err := c.Body()
// 	if err != nil {
// 		slog.Error("Get", slog.Any("err", err))
// 	}

// 	var item any
// 	item = req

// 	res, err := service.NewOps[T](t.s).Get(ctx, item.(models.Item[T]))
// 	if err != nil {
// 		return nil, handleError(err)
// 	}
// 	return res, nil
// }

// func handleError(err error) error {
// 	switch {
// 	case errors.Is(err, internal.ErrNotFound):
// 		return fuego.HTTPError{Err: err, Status: 404, Detail: "Not Found"}
// 	default:
// 		return fuego.HTTPError{Err: err, Status: 500, Detail: "Internal Server Error"}
// 	}
// }
