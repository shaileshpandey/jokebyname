package entities

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
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
		Convey("Testing GetIndividual with invalid Url will error", func() {
			inv := Individual{}

			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(fmt.Sprintf("%s", r), ShouldEqual, "inavlid Url passed to fetch individual")
			}()
			inv.self("@	123")
		})
		Convey("Testing GetIndividual with valid Url will provide correct data", func() {
			inv := Individual{}
			httpmock.Activate()
			httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles", httpmock.NewStringResponder(200, `{"name":"Mikołaj","surname":"Wróbel","gender":"male","region":"Poland"}`))
			httpmock.GetCallCountInfo()
			inv.self("https://api.mybiz.com/articles")
			So(inv.Name, ShouldEqual, "Mikołaj")
			So(inv.SurName, ShouldEqual, "Wróbel")
			So(inv.Gender, ShouldEqual, "male")
			So(inv.Region, ShouldEqual, "Poland")
		})
	})
}
