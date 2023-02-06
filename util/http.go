package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Consistent error return for all failures
type ErrorReturn struct {
	Fail bool   `json:"fail"`
	Msg  string `json:"msg"`
}

// HTTP handler func for client-facing errors
func Reject(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, &ErrorReturn{
		Fail: true,
		Msg:  err.Error(),
	})
}

// HTTP handler func for JSON
func JSON(c echo.Context, x any, errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, x)
}

// Fires when error is not explictly handled by Reject
// Send generic message instead of non-client-facing message
func CustomJsonErrorHandler(err error, c echo.Context) {
	_ = c.JSON(http.StatusInternalServerError, &ErrorReturn{
		Fail: true,
		Msg:  "Internal Server Error",
	})
}
