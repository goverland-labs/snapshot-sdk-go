package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"

	"github.com/goverland-labs/snapshot-sdk-go/client"
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
			ListVotesWithProposalIDsFilter("0x863f2b77141da578251a0dc1e49a4caccf68aa61fb2ecac1249764983748b05f"),
			ListVotesWithPagination(10, 0),
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("vote by vote ID", t, func() {
		vote, err := sdk.VoteByID(
			context.Background(),
			"0xf6c08b692d4bd1fd4c091431895309d38f3d6d7770c075c14b63fca27662b073",
		)

		convey.So(err, convey.ShouldBeNil)
		convey.So(vote, convey.ShouldNotBeNil)
		convey.So(vote.ID, convey.ShouldEqual, "0xf6c08b692d4bd1fd4c091431895309d38f3d6d7770c075c14b63fca27662b073")
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
