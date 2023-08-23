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
			ListProposalWithSpacesFilter("stgdao.eth", "safe.eth"),
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

	convey.Convey("list proposals by IDs", t, func() {
		list, err := sdk.ListProposal(
			context.Background(),
			ListProposalWithIDFilter(
				"QmRS9FH4qBCAt7LaUrTGcbiLGo9zaSzRw5NnUdcwaMKWVo",
				"QmRwFUM2or1cQPK4Vm3RPqHmpuCU4v2LwfAg2NW9emwtSx",
				"QmUsZPqoNAPcoZroumEDWU5vyub5rnWk24KrE2dP9SgDSw",
				"QmRvnAHqcLWyhuKiYw8zbYdhtFYzREmxhfRRvne3ZHgoww",
			),
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 4)
	})
}
