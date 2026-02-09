package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingStatements(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list statements", t, func() {
		list, err := sdk.ListStatements(context.Background(), "paraswap-dao.eth", []string{
			"0x0edEFA91e99da1eDDD1372c1743A63B1595fC413",
			"0x3070f20f86fDa706Ac380F5060D256028a46eC29",
			"0xF38Ad5096E888E0662D59b318371c28eFbaA4181",
		})

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
	})
}
