package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingProposals(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list proposals", t, func() {
		list, err := sdk.ListProposal(context.Background(), ListProposalWithPagination(10, 0))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("list proposals by space IDs", t, func() {
		list, err := sdk.ListProposal(
			context.Background(),
			ListProposalWithSpaces("stgdao.eth", "safe.eth"),
			ListProposalWithPagination(10, 0),
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("get proposal by id", t, func() {
		proposal, err := sdk.GetProposalByID(context.Background(), "0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")

		convey.So(err, convey.ShouldBeNil)
		convey.So(proposal, convey.ShouldNotBeNil)
		convey.So(proposal.ID, convey.ShouldEqual, "0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")
	})
}
