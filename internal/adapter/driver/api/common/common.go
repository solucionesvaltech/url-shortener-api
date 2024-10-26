package common

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"url-shortener/pkg/customerror"
)

func Deserialize(ctx echo.Context, output any) error {
	defer ctx.Request().Body.Close()

	raw, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(raw, output)

	return err
}

func HandleError(
	serverCtx echo.Context,
	code int,
	errorType customerror.ErrorType,
	domain string,
	message string,
	callerError error,
) *echo.HTTPError {
	err := new(customerror.CustomError)

	if errors.As(callerError, &err) {
		serverCtx.Set("customError", err)
		return echo.NewHTTPError(code, err.Error())
	}

	customError := customerror.NewError(errorType, domain, message, callerError.Error())
	serverCtx.Set("customError", customError)
	return echo.NewHTTPError(code, customError.Error())
}

func GetErrorFromContext(ctx echo.Context) *customerror.CustomError {
	var customError *customerror.CustomError
	_ = ConvertInterfaceToStruct(ctx.Get("customError"), &customError)
	return customError
}

func ConvertInterfaceToStruct(input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, output)
}

func DeserializingError(serverCtx echo.Context, err error, domain string) error {
	return HandleError(
		serverCtx,
		http.StatusBadRequest,
		customerror.Mapping,
		domain,
		"an error occur while deserializing",
		err,
	)
}

func InternalError(serverCtx echo.Context, err error, domain string) error {
	return HandleError(
		serverCtx,
		http.StatusInternalServerError,
		customerror.Processing,
		domain,
		"an error occur while processing",
		err,
	)
}

func NotFoundError(serverCtx echo.Context, err error, domain string) error {
	return HandleError(
		serverCtx,
		http.StatusNotFound,
		customerror.Processing,
		domain,
		"an error occur while processing",
		err,
	)
}
