package client

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Yamashou/gqlgenc/graphqljson"
	"github.com/smartystreets/goconvey/convey"

	"github.com/goverland-labs/snapshot-sdk-go/helpers"
)

const snapshotHub = "https://hub.snapshot.org/graphql"

func TestClient(t *testing.T) {
	c := NewClient(http.DefaultClient, snapshotHub, nil)

	convey.Convey("ListRanking", t, func() {
		list, err := c.ListRanking(context.Background(), 0, 20, "grant", "1")

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetRanking().GetItems(), convey.ShouldHaveLength, 20)
	})

	convey.Convey("ListSpaces", t, func() {
		list, err := c.ListSpaces(context.Background(), 0, 2, nil)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetSpaces(), convey.ShouldHaveLength, 2)
	})

	convey.Convey("GetSpaceByID", t, func() {
		list, err := c.SpaceByID(context.Background(), "stgdao.eth")

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetSpaces(), convey.ShouldHaveLength, 1)
		convey.So(list.GetSpaces()[0].GetID(), convey.ShouldEqual, "stgdao.eth")
	})

	convey.Convey("ListProposals", t, func() {
		list, err := c.ListProposals(context.Background(), 0, 10, 0, nil, nil, "created", OrderDirectionDesc)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetProposals(), convey.ShouldHaveLength, 10)
	})

	convey.Convey("ListProposals filtered by spaces", t, func() {
		list, err := c.ListProposals(context.Background(), 0, 10, 0, []*string{helpers.Ptr("stgdao.eth")}, nil, "created", OrderDirectionDesc)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetProposals(), convey.ShouldHaveLength, 10)
		convey.So(list.GetProposals()[0].Space.GetID(), convey.ShouldEqual, "stgdao.eth")
	})

	convey.Convey("GetProposal by ID", t, func() {
		list, err := c.ProposalByID(context.Background(), "0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetProposals(), convey.ShouldHaveLength, 1)
		convey.So(list.GetProposals()[0].GetID(), convey.ShouldEqual, "0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")
	})

	convey.Convey("ListVotes", t, func() {
		list, err := c.ListVotes(context.Background(), []*string{helpers.Ptr("0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")}, 0, 10, "created", OrderDirectionDesc, 0)

		convey.So(err, convey.ShouldBeNil)
		convey.So(list.GetVotes(), convey.ShouldHaveLength, 10)
		convey.So(list.GetVotes()[0].GetProposal().GetID(), convey.ShouldEqual, "0x108a9e597560c4f249cd8be23acd409059fcd17bb2290d69a550ac2232676e7d")
	})
}

func TestAnyUnmarshal(t *testing.T) {
	convey.Convey("unmarshal any field to map[string]interface{}", t, func() {
		const d = `[{"name":"eth-balance","network":"1","params":{"symbol":"ETH"}}]`
		var strategies []StrategyFragment
		err := graphqljson.UnmarshalData(json.RawMessage(d), &strategies)

		convey.So(err, convey.ShouldBeNil)
		convey.So(strategies, convey.ShouldHaveLength, 1)
		convey.So(strategies[0].Name, convey.ShouldEqual, "eth-balance")
		convey.So(strategies[0].Network, convey.ShouldNotBeNil)
		convey.So(*strategies[0].Network, convey.ShouldEqual, "1")
		convey.So(strategies[0].Params["symbol"], convey.ShouldEqual, "ETH")
	})
}
