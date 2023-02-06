package server

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"

	v "github.com/jellydator/validation"
	// "github.com/jellydator/validation/is"

	"github.com/kofalt/example-ent-crdb/data"
)

func Paginate(c echo.Context, defaultLimit int) (*data.Paginate, error) {

	oS := c.QueryParam("offset")
	lS := c.QueryParam("limit")

	if oS == "" {
		oS = "0"
	}
	if lS == "" {
		lS = fmt.Sprint(defaultLimit)
	}

	o, err1 := strconv.ParseInt(oS, 10, 0)
	l, err2 := strconv.ParseInt(lS, 10, 0)

	if err1 != nil && err2 != nil {
		return nil, errors.New("Limit and offset must be integers")
	}

	p := &data.Paginate{Offset: int(o), Limit: int(l)}

	err := v.ValidateStruct(p,
		v.Field(&p.Limit, v.Min(0), v.Max(200)),
		v.Field(&p.Offset, v.Min(0)),
	)
	if err != nil {
		return nil, err
	}

	return p, nil
}
