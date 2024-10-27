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
	s       *service.Service
	subtype string
}

func NewOps[T models.DataType](s *service.Service, subtype string) Op[T] {
	return Op[T]{s: s, subtype: subtype}
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

// func (t Op[T]) Get(c *gin.Context, req *Req) (*Res, error) {
// 	ctx, reqModel := t.logSubtaskReq(c, *req)
// 	// txID := c.Value("transactionID").(string)
// 	// _, span := t.s.Otel.StartSpan(ctx, "Get Task")
// 	// defer span.End()

// 	// span.Set("transactionID", txID)
// 	// span.Set("taskType", t.subtype)
// 	var r interface{}
// 	r = reqModel
// 	res, err := service.NewOps[T](t.s).Get(ctx, r.(models.Item[T]))
// 	if err != nil {
// 		// span.Error(err)
// 		return nil, handleTaskError(err)
// 	}
// 	return t.formatResponse(res), nil
// }

// func (t Op[T, Req, Res]) Create(c *gin.Context, req *Req) (*Res, error) {
// 	ctx, reqModel := t.logSubtaskReq(c, *req)
// 	// txID := c.Value("transactionID").(string)
// 	// ctx, span := t.s.Otel.StartSpan(ctx, "Create Task")
// 	// defer span.End()

// 	// span.Set("transactionID", txID)
// 	// span.Set("taskType", t.subtype)
// 	var r interface{}
// 	r = reqModel
// 	res, err := service.NewOps[T](t.s).Create(ctx, r.(models.Item[T]))
// 	if err != nil {
// 		// span.Error(err)
// 		return nil, handleTaskError(err)
// 	}

// 	return t.formatResponse(res), nil
// }

// func (t Op[T, Req, Res]) Patch(c *gin.Context, req *Req) (*Res, error) {
// 	ctx, reqModel := t.logSubtaskReq(c, *req)
// 	// txID := c.Value("txID").(string)
// 	// ctx, span := t.s.Otel.StartSpan(ctx, "Patch Task")
// 	// defer span.End()

// 	// span.Set("transactionID", txID)
// 	// span.Set("taskType", t.subtype)
// 	var r interface{}
// 	r = reqModel
// 	res, err := service.NewOps[T](t.s).Patch(ctx, r.(models.Item[T]))
// 	if err != nil {
// 		// span.Error(err)
// 		return nil, handleTaskError(err)
// 	}
// 	return t.formatResponse(res), nil
// }

// func (t Op[T, Req, Res]) Delete(c *gin.Context, req *Req) error {
// 	ctx, reqModel := t.logSubtaskReq(c, *req)
// 	// txID := c.Value("txID").(string)
// 	// ctx, span := t.s.Otel.StartSpan(ctx, "Delete Task")
// 	// defer span.End()

// 	// span.Set("transactionID", txID)
// 	// span.Set("taskType", t.subtype)
// 	var r interface{}
// 	r = reqModel
// 	err := service.NewOps[T](t.s).Delete(ctx, r.(models.Item[T]))
// 	if err != nil {
// 		// span.Error(err)
// 		return handleTaskError(err)
// 	}

// 	return nil
// }

// func (t Op[T, Req, Res]) formatResponse(data *T) *Res {
// 	var r interface{}
// 	r = data
// 	item := r.(models.ResponseItem[T, Res])
// 	return item.ToResponse()
// }

// func handleTaskError(err error) error {
// 	slogctx.Error(nil, "handleTaskError", slog.Any("err", err))
// 	if strings.Contains(err.Error(), "there is no unique or exclusion constraint matching the ON CONFLICT specification") {
// 		return internal.ErrBadRequest
// 	}
// 	return err
// }

// func (t Op[T, Req, Res]) logSubtaskReq(c *gin.Context, req Req) (context.Context, *T) {
// 	return logTaskCtx(c), req.ToModel(t.subtype)
// }

// func logTaskCtx(c *gin.Context) context.Context {
// 	txID := c.Param("txID")
// 	ctx := logTxReq(c, &txID)
// 	parts := strings.Split(c.FullPath(), "/")
// 	if len(parts) > 0 {
// 		last := parts[len(parts)-1]
// 		if last != "" && last != "task" {
// 			ctx = slogctx.WithValue(ctx, "taskType", parts[len(parts)-1])
// 		}
// 	}

// 	return ctx
// }
