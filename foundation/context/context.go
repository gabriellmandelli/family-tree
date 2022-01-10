package util

import (
	"context"

	"github.com/gabriellmandelli/family-tree/foundation/model"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type loggerKeyType string

func getContext(c echo.Context) context.Context {
	return c.Request().Context()
}

func InitializeContext(echoCtx echo.Context, body interface{}) (context.Context, *model.RequestData, *errorx.Error) {
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

	return ctx, &requestData, errx
}
