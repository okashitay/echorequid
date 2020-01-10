package echorequid

import (
	"bytes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"net"
)

// func CalcRemoteAddr(req engine.Request) string
// calcurate remote address from request
func CalcRemoteAddr(req engine.Request) string {
	ra := req.RemoteAddress()
	if ip := req.Header().Get(echo.HeaderXRealIP); ip != "" {
		ra = ip
	} else if ip = req.Header().Get(echo.HeaderXForwardedFor); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}

	return ra
}

// func BuildString(strs ...string) string
// concat all strings
func BuildString(strs ...string) string {
	var buffer bytes.Buffer
	for _, s := range strs {
		buffer.WriteString(s)
	}
	return buffer.String()
}
