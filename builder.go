package echorequid

import "github.com/labstack/echo"

// RequestUUIDBuilder interface
type RequestUUIDBuilder interface {
	Build(c echo.Context) (string, error)
}
