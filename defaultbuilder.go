package echorequid

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"time"
)

// DefaultBuilder struct
type DefaultBuilder struct{}

// func (d DefaultBuilder) Build(c echo.Context) (string, error)
// Build UUID using remoteaddr, and time.Now()
func (d DefaultBuilder) Build(c echo.Context) (string, error) {
	base := BuildString(CalcRemoteAddr(c.Request()), time.Now().String())
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf(base)))), nil
}
