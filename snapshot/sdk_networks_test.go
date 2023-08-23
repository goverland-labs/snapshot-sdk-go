package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingNetworks(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list networks", t, func() {
		list, err := sdk.ListNetworks(context.Background())

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
	})
}
