package echorequid

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCalcRemoteAddr(t *testing.T) {
	Convey("Return X-Real-IP", t, func() {
		expected := "17.151.0.151"
		dummy := "172.151.0.5"
		req := test.NewRequest(echo.GET, "/", nil)
		req.Header().Set(echo.HeaderXRealIP, expected)
		req.Header().Set(echo.HeaderXForwardedFor, dummy)

		actual := CalcRemoteAddr(req)
		So(actual, ShouldEqual, expected)
	})

	Convey("Return X-Forwarded-For", t, func() {
		expected := "172.151.0.5"
		dummy := ""
		req := test.NewRequest(echo.GET, "/", nil)
		req.Header().Set(echo.HeaderXRealIP, dummy)
		req.Header().Set(echo.HeaderXForwardedFor, expected)

		actual := CalcRemoteAddr(req)
		So(actual, ShouldEqual, expected)
	})
}
