package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingMessages(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list messages", t, func() {
		list, err := sdk.ListMessage(context.Background(), ListMessageWithPagination(10, 0))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("list messages with skipping mci", t, func() {
		list, err := sdk.ListMessage(context.Background(), ListMessageWithPagination(10, 0), ListMessageWithMCIFilter(1000))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
		for i, msg := range list {
			mci := msg.GetMci()

			convey.So(mci, convey.ShouldNotBeNil)
			convey.So(*mci, convey.ShouldEqual, 1000+i+1)
		}
	})
}
