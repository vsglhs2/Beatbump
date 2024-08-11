package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var Logger zerolog.Logger = NewLogger()

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// call the next middleware/handler
		err := next(c)

		if err != nil {
			Logger.Error().Fields(map[string]interface{}{
				"query":    c.Request().URL.RawQuery,
				"response": c.Response().Status,
				"error":    err,
			}).Msg(c.Request().URL.Path)
			return err
		} else {
			Logger.Info().Fields(map[string]interface{}{
				"query":    c.Request().URL.RawQuery,
				"response": c.Response().Status,
				"UA":       c.Request().UserAgent(),
			}).Msg(c.Request().URL.Path)
		}

		return nil
	}
}

func NewLogger() zerolog.Logger {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	Logger := zerolog.New(output).With().Caller().Timestamp().Logger()

	return Logger
}
