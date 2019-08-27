package entities

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGetJoke(t *testing.T) {
	Convey("Testing GetJoke", t, func() {
		Convey("Testing GetJoke with blank Url will error", func() {
			inv := Joke{}

			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(fmt.Sprintf("%s", r), ShouldEqual, "Unable to fetch random joke")
			}()
			inv.self("")
		})
		Convey("Testing GetJoke with invalid Url will error", func() {
			inv := Joke{}

			defer func() {
				r := recover()
				So(r, ShouldNotBeNil)
				So(fmt.Sprintf("%s", r), ShouldEqual, "inavlid Url passed to fetch joke")
			}()
			inv.self("@	123")
		})
		Convey("Testing GetJoke with valid Url will provide correct data", func() {
			inv := Joke{}
			httpmock.Activate()
			httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles", httpmock.NewStringResponder(200, `{ "type": "success", "value": { "id": 478, "joke": "John Doe can instantiate an abstract class.", "categories": ["nerdy"] } }`))
			httpmock.GetCallCountInfo()
			inv.self("https://api.mybiz.com/articles")
			So(inv.Type, ShouldEqual, "success")
			So(inv.Value.Text, ShouldEqual, "John Doe can instantiate an abstract class.")

		})
	})
}
