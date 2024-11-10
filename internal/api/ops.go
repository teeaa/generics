package api

import (
	"log/slog"

	"github.com/go-fuego/fuego"
	"github.com/teeaa/generics/internal/models"
)

func Create[T models.DataType](c *fuego.ContextWithBody[T]) (*T, error) {
	req, err := c.Body() // req is of type T
	if err != nil {
		slog.Error("Create", slog.Any("err", err))
	}

	slog.Info("Create", slog.Any("req", req))

	// id := uuid.New().String()
	// txID := c.PathParam("txid")
	// req.ID = id

	return nil, nil
}
