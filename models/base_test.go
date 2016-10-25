package models_test

import (
	"net/http"
	"testing"

	"github.com/apiseedprojects/go/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProductModelBase(t *testing.T) {
	Convey("Given a base framewrok", t, func() {
		Convey("When I create a new ModelError", func() {
			me := models.NewModelError(http.StatusInternalServerError, "my message: %s", "argument")
			Convey("Then I should get a new ModelError", func() {
				So(me, ShouldNotBeNil)
			})
			Convey("And I should be able to call its Error() function", func() {
				So(me.Error(), ShouldEqual, "my message: argument")
			})
			Convey("And it should be an instance of error interface", func() {
				meint := interface{}(me)
				_, ok := (meint).(error)
				So(ok, ShouldBeTrue)
			})
		})
	})
}
