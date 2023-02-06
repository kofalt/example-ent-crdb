package util

import (
	"bytes"
	"os"

	"github.com/kofalt/echotozero"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// NewGlobalZero configures an instanced and global logger
// Returns logger and logging level string.
//
// When not in JSON mode, fields named "ex" will be pretty-printed.
func NewGlobalZero(jsonFormat bool, verbosity int) (zerolog.Logger, string) {
	var zLog zerolog.Logger

	if jsonFormat {
		zLog = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		zLog = zerolog.New(zerolog.ConsoleWriter{
			Out: os.Stdout,

			FieldsExclude: []string{
				// Skip debug field name from normal print
				"ex",

				// Skip request ID in debug mode
				"rid",
			},

			// Send debug field name in debug JSON format
			FormatExtra: func(evt map[string]interface{}, buf *bytes.Buffer) error {
				ex, ok := evt["ex"]
				if ok {
					buf.WriteString("\n" + Debug(ex))
				}
				return nil
			},
		}).With().Timestamp().Logger()
	}

	log.Logger = zLog

	lvl := zerolog.Level(verbosity)
	zerolog.SetGlobalLevel(lvl)
	zerolog.DurationFieldInteger = true

	return zLog, lvl.String()
}

// NewEcho configures a new API server with message logging but no middleware
func NewEcho(zLog zerolog.Logger) (*echo.Echo, *echotozero.Logger) {
	e := echo.New()

	adapter := echotozero.New(zLog)

	e.HideBanner = true
	e.Logger = adapter
	e.HTTPErrorHandler = CustomJsonErrorHandler

	return e, adapter
}
