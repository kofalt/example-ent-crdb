package server

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	// "github.com/rs/zerolog/log"

	v "github.com/jellydator/validation"
	// "github.com/jellydator/validation/is"

	"github.com/kofalt/example-ent-crdb/data"
	"github.com/kofalt/example-ent-crdb/ent"
	. "github.com/kofalt/example-ent-crdb/util"
)

func (s *Server) GetUsers(c echo.Context) error {
	p, err := Paginate(c, 100)
	if err != nil {
		return Reject(c, err)
	}

	r, err := s.Db.GetUsers(&data.UserSearch{
		Paginate: p,
		City:     c.QueryParam("city"),
	})

	return JSON(c, r, err)
}

func (s *Server) GetUser(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return Reject(c, err)
	}

	r, err := s.Db.GetUser(id)

	return JSON(c, r, err)
}

func (s *Server) SetUserAddress(c echo.Context) error {

	user := new(ent.User)
	err := c.Bind(user)
	if err != nil {
		return Reject(c, err)
	}

	err = v.ValidateStruct(user,
		v.Field(&user.Address, v.Length(1, 1000)),
	)
	if err != nil {
		return err
	}

	r, err := s.Db.MoveUser(user)

	return JSON(c, r, err)
}

func (s *Server) TopRides(c echo.Context) error {

	r, err := s.Db.TopRides()

	if err != nil {
		return Reject(c, err)
	}
	return c.JSON(http.StatusOK, r)
}

func (s *Server) Version(c echo.Context) error {
	return c.JSON(http.StatusOK, MAP{
		"BuildHash": s.BuildHash,
		"BuildDate": s.BuildDate,
	})
}
