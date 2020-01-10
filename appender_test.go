package echorequid

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"./mock"
)

// func TestAppendRequestUUIDWithConfigAndBuilder(t *testing.T)
// test func AppendRequestUUIDWithConfigAndBuilder
// see appender.go
func TestAppendRequestUUIDWithConfigAndBuilder(t *testing.T) {
	// setup mock
	headername := "X-Request-Id"
	mc := gomock.NewController(t)
	defer mc.Finish()
	h := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	Convey("Do not overwrite when configured to PASSTHROUGH strategy", t, func() {
		expected := "THIS_WILL_NOT_BE_OVERWRITTEN"
		dummy := "WILL_FAIL_IF_RETURNED_THIS"
		cfg := &AppendRequestUUIDConfig{
			HeaderFieldName:   headername,
			DuplicateStrategy: PASSTHROUGH,
		}
		e := echo.New()
		req := test.NewRequest(echo.GET, "/", nil)
		req.Header().Set(headername, expected)
		rec := test.NewResponseRecorder()
		c := e.NewContext(req, rec)
		h := func(c echo.Context) error {
			return c.String(http.StatusOK, "test")
		}

		mb := mock_echorequid.NewMockRequestUUIDBuilder(mc)
		mb.EXPECT().Build(gomock.Any()).AnyTimes().Return(dummy, nil)

		err := AppendRequestUUIDWithConfigAndBuilder(cfg, mb)(h)(c)
		So(err, ShouldBeNil)
		So(c.Request().Header().Get(headername), ShouldEqual, expected)
	})
	Convey("Overwrite when configured to OVERWRITE strategy", t, func() {
		dummy := "THIS_WILL_BE_OVERWRITTEN"
		expected := "WILL_SUCCESS_IF_RETURNED_THIS"
		cfg := &AppendRequestUUIDConfig{
			HeaderFieldName:   headername,
			DuplicateStrategy: OVERWRITE,
		}
		e := echo.New()
		req := test.NewRequest(echo.GET, "/", nil)
		req.Header().Set(headername, dummy)
		rec := test.NewResponseRecorder()
		c := e.NewContext(req, rec)
		mb := mock_echorequid.NewMockRequestUUIDBuilder(mc)
		mb.EXPECT().Build(gomock.Any()).AnyTimes().Return(expected, nil)

		err := AppendRequestUUIDWithConfigAndBuilder(cfg, mb)(h)(c)
		So(err, ShouldBeNil)
		So(c.Request().Header().Get(headername), ShouldEqual, expected)
	})

	Convey("Confirm specific call", t, func() {

		e := echo.New()
		req := test.NewRequest(echo.GET, "/", nil)
		rec := test.NewResponseRecorder()
		c := e.NewContext(req, rec)

		err := AppendRequestUUID()(h)(c)
		So(err, ShouldBeNil)
		So(c.Request().Header().Get(headername), ShouldNotBeNil)

		cfg := &AppendRequestUUIDConfig{
			HeaderFieldName:   headername,
			DuplicateStrategy: OVERWRITE,
		}
		err = AppendRequestUUIDWithConfig(cfg)(h)(c)
		So(err, ShouldBeNil)
		So(c.Request().Header().Get(headername), ShouldNotBeNil)

		mb := mock_echorequid.NewMockRequestUUIDBuilder(mc)
		mb.EXPECT().Build(gomock.Any()).AnyTimes().Return("", errors.New("test-error"))
		req.Header().Del(headername)
		err = AppendRequestUUIDWithBuilder(mb)(h)(c)
		So(err, ShouldNotBeNil)
	})
}
