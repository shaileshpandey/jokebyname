package entities

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestGetIndividualwithjoke(t *testing.T) {
	Convey("Testing GetIndividualwithjoke", t, func() {
		Convey("Testing GetIndividualwithjoke with blank Url will error", func() {
			inv := IndividualWithJoke{}
			out := inv.TellRandomJoke("", "https://www.google.com")
			So(out, ShouldEqual, "Unable to fetch random person")
		})
		Convey("Testing GetIndividualwithjoke with invalid Url will error", func() {
			inv := IndividualWithJoke{}
			out := inv.TellRandomJoke("https://www.google.com", "")
			So(out, ShouldEqual, "Unable to fetch random person")
		})
		Convey("Testing GetIndividualwithjoke with valid Url will provide correct data", func() {
			inv := IndividualWithJoke{}
			httpmock.Activate()
			httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles", httpmock.NewStringResponder(200, `{"name":"Mikołaj","surname":"Wróbel","gender":"male","region":"Poland"}`))
			httpmock.RegisterResponder("GET", "https://www.google.com", httpmock.NewStringResponder(200, `{ "type": "success", "value": { "id": 478, "joke": "John Doe can instantiate an abstract class.", "categories": ["nerdy"] } }`))
			httpmock.GetCallCountInfo()
			httpmock.GetCallCountInfo()
			inv.TellRandomJoke("https://api.mybiz.com/articles", "https://www.google.com")
			So(inv.CurrentJoke.Type, ShouldEqual, "success")
			So(inv.CurrentJoke.Value.Text, ShouldEqual, "John Doe can instantiate an abstract class.")
		})
	})
}
