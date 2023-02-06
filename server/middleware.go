package server

import (
	"strconv"
	"time"

	"github.com/jaevor/go-nanoid"
	"github.com/kofalt/echotozero"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// MiddleWare logs successful requests to debug, failed req to error, and never skips.
func Middleware(logger *echotozero.Logger) echo.MiddlewareFunc {
	return MiddlewareWithOptions(
		logger,
		zerolog.DebugLevel,
		zerolog.ErrorLevel,
		middleware.DefaultSkipper,
	)
}

func MiddlewareWithOptions(logger *echotozero.Logger, successLvl, failLvl zerolog.Level, skipper middleware.Skipper) echo.MiddlewareFunc {

	randomIDGenerator, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			// Set rid header
			rid := randomIDGenerator()
			c.Response().Header().Add("X-Request-ID", rid)

			// Allows some requests to skip logging
			if skipper(c) {
				return next(c)
			}

			start := time.Now()
			req := c.Request()
			res := c.Response()

			// Invokes next middleware or handler
			err := next(c)

			var evt *zerolog.Event
			if err != nil {
				evt = logger.ZLog.WithLevel(failLvl).Err(err)
			} else {
				evt = logger.ZLog.WithLevel(successLvl)
			}

			cl := req.Header.Get(echo.HeaderContentLength)
			if cl == "" {
				cl = "0"
			}

			elapsed := time.Since(start)
			evt.Dur("ms", elapsed)
			// evt.Str("host", req.Host)
			// evt.Str("ip", c.RealIP())
			evt.Str("reqB", cl)
			evt.Str("resB", strconv.FormatInt(res.Size, 10))
			evt.Int("resC", res.Status)
			evt.Str("uri", req.RequestURI)
			// evt.Str("agent", req.UserAgent())

			evt.Str("rid", rid)

			evt.Send()
			return err
		}
	}
}
