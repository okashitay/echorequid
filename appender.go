package echorequid

import (
	"github.com/labstack/echo"
)

type Strategy int

const (
	OVERWRITE Strategy = iota
	PASSTHROUGH
)

// type AppendRequestUUIDConfig struct
type AppendRequestUUIDConfig struct {
	HeaderFieldName   string
	DuplicateStrategy Strategy
}

var DefaultAppendRequestUUIDConfig = &AppendRequestUUIDConfig{
	HeaderFieldName:   "X-Request-Id",
	DuplicateStrategy: PASSTHROUGH,
}

// func AppendRequestUUID() echo.MiddlewareFunc
// Generate Middleware func for echo
// Using DefaultConfiguration and  DefaultBUilder
func AppendRequestUUID() echo.MiddlewareFunc {
	return AppendRequestUUIDWithConfigAndBuilder(DefaultAppendRequestUUIDConfig, DefaultBuilder{})
}

// func AppendRequestUUIDWithBuilder(builder RequestUUIDBuilder) echo.MiddlewareFunc
// Generate Middleware func for echo
// With builder implementation Using Default Config
func AppendRequestUUIDWithBuilder(builder RequestUUIDBuilder) echo.MiddlewareFunc {
	return AppendRequestUUIDWithConfigAndBuilder(DefaultAppendRequestUUIDConfig, builder)
}

// Generate Middleware func for echo
// With configuration Using DefaultBUilder
func AppendRequestUUIDWithConfig(config *AppendRequestUUIDConfig) echo.MiddlewareFunc {
	return AppendRequestUUIDWithConfigAndBuilder(config, DefaultBuilder{})
}

// func AppendRequestUUIDWithConfigAndBuilder(config *AppendRequestUUIDConfig, builder RequestUUIDBuilder) echo.MiddlewareFunc
// Generate Middleware func for echo
// Setting calcurated RequesUUID to HTTP Request Header,
// With configuration and builder implementation
func AppendRequestUUIDWithConfigAndBuilder(config *AppendRequestUUIDConfig, builder RequestUUIDBuilder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// if choice PASSTHROUGH and header contains key, use it
			if config.DuplicateStrategy == PASSTHROUGH &&
				c.Request().Header().Contains(config.HeaderFieldName) {
				return next(c)
			}
			id, err := builder.Build(c)
			if err != nil {
				return err
			}
			c.Request().Header().Set(config.HeaderFieldName, id)
			return next(c)
		}
	}
}
