package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func StatsEndpointHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, ([]byte)("version:"+time.Now().String()))
}
