package entities

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetIndividual(t *testing.T) {
	Convey("Testing GetIndividual", t, func() {
		Convey("Testing GetIndividual with blank Url will error", func() {
			inv := Individual{}

			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(fmt.Sprintf("%s", r), ShouldEqual, "Unable to fetch random person")
			}()
			inv.self("")
		})
	})
}
