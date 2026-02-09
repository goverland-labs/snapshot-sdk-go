package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingUsers(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("user about info", t, func() {
		list, err := sdk.ListUsers(context.Background(), []string{
			"0x06aD51E6CC8AEaFcc0aAf2df6Dc7870d365dd8E1",
			"0xd5D171a9AA125AF13216C3213B5A9Fc793FcCF2c",
			"0x117B4B046ad3B7f152A688dc9E5461c53B512dAC",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
	})
}
