package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSDK_Validate(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("fail validate for vote", t, func() {
		validate, err := sdk.Validate(context.Background(), ValidationParams{
			Validation: "passport-gated",
			Author:     "0x7697cAB0e123c68d27d7D5A9EbA346d7584Af888",
			Space:      "nwnwert.eth",
			Network:    "1",
			Snapshot:   18698857,
			Params: map[string]interface{}{
				"scoreThreshold": 14,
				"operator":       "NONE",
			},
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(validate, convey.ShouldNotBeNil)

		convey.So(validate.Result, convey.ShouldBeFalse)
	})
}

func TestSDK_GetVotingPower(t *testing.T) {
	sdk := NewSDK()

	// This test for manual checking only because it's based on the proposal in Snapshot's database
	convey.SkipConvey("success get voting power", t, func() {
		vp, err := sdk.GetVotingPower(context.Background(), GetVotingPowerParams{
			Voter:    "0x7697cAB0e123c68d27d7D5A9EbA346d7584Af888",
			Space:    "nwnwert.eth",
			Proposal: "0xfc37b379c65cbef382276030741fa5db62c0328eb0d6cb51ac3f187f0e75448c",
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(vp, convey.ShouldNotBeNil)

		convey.So(*vp.GetVp(), convey.ShouldEqual, 1)
	})

	convey.Convey("get voting power with invalid address", t, func() {
		_, err := sdk.GetVotingPower(context.Background(), GetVotingPowerParams{
			Voter:    "fake",
			Space:    "nwnwert.eth",
			Proposal: "0xfc37b379c65cbef382276030741fa5db62c0328eb0d6cb51ac3f187f0e75448c",
		})
		convey.So(err, convey.ShouldNotBeNil)
	})
}
