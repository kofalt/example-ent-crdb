package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/kofalt/example-ent-crdb/data"
	"github.com/kofalt/example-ent-crdb/util"
)

type Config struct {
	util.DbConnectSettings

	LogJSON      bool
	LogVerbosity int

	HttpPort string

	BuildHash string
	BuildDate string
}

type Server struct {
	*echo.Echo

	*Config

	Db *data.DB
}

func (s *Server) Serve() {
	err := s.Start(":" + s.HttpPort)

	log.Fatal().Err(err).Send()
}

// Bootstrap calls various service allocators
func Bootstrap(c *Config) (*Server, error) {
	// Logging before this line will be unconfigured
	zLog, logLevel := util.NewGlobalZero(c.LogJSON, c.LogVerbosity)

	log.Info().
		Str("lvl", logLevel).
		Str("db", c.DbName+"@"+c.DbHost+":"+c.DbPort+" "+c.DbUser).
		Msg("Config")

	db, err := data.Setup(&c.DbConnectSettings)
	if err != nil {
		return nil, err
	}

	// Service API
	logger := zLog.With().Str("c", "http").Logger()
	echoServer, echoAdapter := util.NewEcho(logger)
	echoServer.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableStackAll: true,
	}))
	echoServer.Use(Middleware(echoAdapter))

	s := &Server{echoServer, c, db}
	s.Routes()
	return s, nil
}
