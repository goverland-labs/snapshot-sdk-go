package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/goverland-labs/sdk-snapshot-go/client"
)

func TestGettingVotes(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list votes", t, func() {
		list, err := sdk.ListVotes(context.Background(), ListVotesWithPagination(10, 0))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("list votes by proposal ID", t, func() {
		list, err := sdk.ListVotes(
			context.Background(),
			ListVotesWithProposalIDsFilter("QmPvbwguLfcVryzBRrbY4Pb9bCtxURagdv1XjhtFLf3wHj"),
			ListVotesWithPagination(10, 0),
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("list votes ordered by created_at", t, func() {
		list, err := sdk.ListVotes(
			context.Background(),
			ListVotesWithOrderBy("created", client.OrderDirectionDesc),
			ListVotesWithPagination(2, 0),
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 2)
		convey.So(list[0].Created >= list[1].Created, convey.ShouldBeTrue)
	})
}
