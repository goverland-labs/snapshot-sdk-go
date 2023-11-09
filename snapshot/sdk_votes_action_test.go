package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/goverland-labs/sdk-snapshot-go/helpers"
)

func TestSDK_Validate(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("fail validate for vote", t, func() {
		validate, err := sdk.Validate(context.Background(), ValidationParams{
			Validation: "passport-gated",
			Author:     "0xB9879cD06c904c2FDbc75d03534929b5E842F3a0",
			Space:      "nwnwert.eth",
			Network:    "1",
			Snapshot:   18480218,
			Params: map[string]interface{}{
				"stamps":   []string{"Gitcoin"},
				"operator": "OR",
			},
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(validate, convey.ShouldNotBeNil)

		convey.So(validate.Result, convey.ShouldBeFalse)
	})

	convey.Convey("validate with invalid address", t, func() {
		_, err := sdk.Validate(context.Background(), ValidationParams{
			Validation: "passport-gated",
			Author:     "fake-address",
			Space:      "nwnwert.eth",
			Network:    "1",
			Snapshot:   18480218,
			Params: map[string]interface{}{
				"stamps":   []string{"Gitcoin"},
				"operator": "OR",
			},
		})
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestSDK_GetVotingPower(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("success get voting power", t, func() {
		vp, err := sdk.GetVotingPower(context.Background(), GetVotingPowerParams{
			Address: "0x7697cAB0e123c68d27d7D5A9EbA346d7584Af888",
			Network: "1",
			Strategies: []StrategyFragment{
				{
					Name:    "eth-balance",
					Network: helpers.Ptr("1"),
					Params:  make(map[string]interface{}),
				},
			},
			Snapshot:   18480218,
			Space:      "nwnwert.eth",
			Delegation: false,
		})
		convey.So(err, convey.ShouldBeNil)
		convey.So(vp, convey.ShouldNotBeNil)

		convey.So(vp.Result.VP, convey.ShouldEqual, 0.013164824760661685)
	})

	convey.Convey("get voting power with invalid address", t, func() {
		_, err := sdk.GetVotingPower(context.Background(), GetVotingPowerParams{
			Address: "fake address",
			Network: "1",
			Strategies: []StrategyFragment{
				{
					Name:    "eth-balance",
					Network: helpers.Ptr("1"),
					Params:  make(map[string]interface{}),
				},
			},
			Snapshot:   18480218,
			Space:      "nwnwert.eth",
			Delegation: false,
		})
		convey.So(err, convey.ShouldNotBeNil)
	})
}
