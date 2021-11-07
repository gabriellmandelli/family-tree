package util

import (
	"context"

	"github.com/gabriellmandelli/family-tree/internal/model"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type loggerKeyType string

const (
	requestDataKey loggerKeyType = "requestDataKey"
)

func getContext(c echo.Context) context.Context {
	return c.Request().Context()
}

func InitializeContext(echoCtx echo.Context, body interface{}) (*context.Context, *model.RequestData, *errorx.Error) {
	var errx *errorx.Error

	ctx := getContext(echoCtx)

	if body != nil {
		if err := echoCtx.Bind(body); err != nil {
			errx := errorx.IllegalArgument.Wrap(err, "Failed to parse request body")
			echoCtx.Logger().Warn(errx)
		}
	}

	requestData := model.RequestData{
		Headers:     echoCtx.Request().Header,
		QueryParams: echoCtx.QueryParams(),
	}

	ctx = context.WithValue(ctx, requestDataKey, requestData)

	return &ctx, &requestData, errx
}
