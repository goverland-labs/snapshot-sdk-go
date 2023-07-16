package snapshot

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGettingSpaces(t *testing.T) {
	sdk := NewSDK()

	convey.Convey("list spaces", t, func() {
		list, err := sdk.ListSpace(context.Background(), ListSpaceWithPagination(10, 0))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 10)
	})

	convey.Convey("list spaces by IDs", t, func() {
		list, err := sdk.ListSpace(context.Background(), ListSpaceWithIDFilter("stgdao.eth", "safe.eth"))

		convey.So(err, convey.ShouldBeNil)
		convey.So(list, convey.ShouldNotBeNil)
		convey.So(list, convey.ShouldHaveLength, 2)

		ids := make([]string, 0, len(list))
		for i := range list {
			ids = append(ids, list[i].ID)
		}

		convey.So(ids, convey.ShouldContain, "stgdao.eth")
		convey.So(ids, convey.ShouldContain, "safe.eth")
	})

	convey.Convey("get space by id", t, func() {
		space, err := sdk.GetSpaceByID(context.Background(), "stgdao.eth")

		convey.So(err, convey.ShouldBeNil)
		convey.So(space, convey.ShouldNotBeNil)
		convey.So(space.ID, convey.ShouldEqual, "stgdao.eth")
	})

	convey.Convey("get ranked list", t, func() {
		ids, err := sdk.GetRanking(context.Background(), RankingWithPagination(20, 0))

		convey.So(err, convey.ShouldBeNil)
		convey.So(ids, convey.ShouldHaveLength, 20)
	})

	convey.Convey("get ranked list by network", t, func() {
		ids, err := sdk.GetRanking(context.Background(), RankingWithNetwork("1"))

		convey.So(err, convey.ShouldBeNil)
		convey.So(ids, convey.ShouldHaveLength, 20)
	})
}
