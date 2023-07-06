// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package client

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type SnapshotClient interface {
	ListNetworks(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*ListNetworks, error)
	ListProposals(ctx context.Context, skip int64, first int64, createdAfter int64, spaces []*string, orderBy string, orderDirection OrderDirection, interceptors ...clientv2.RequestInterceptor) (*ListProposals, error)
	ProposalByID(ctx context.Context, id string, interceptors ...clientv2.RequestInterceptor) (*ProposalByID, error)
	ListRanking(ctx context.Context, skip int64, first int64, category string, network string, interceptors ...clientv2.RequestInterceptor) (*ListRanking, error)
	ListSpaces(ctx context.Context, skip int64, first int64, ids []*string, interceptors ...clientv2.RequestInterceptor) (*ListSpaces, error)
	SpaceByID(ctx context.Context, id string, interceptors ...clientv2.RequestInterceptor) (*SpaceByID, error)
	ListVotes(ctx context.Context, proposal string, skip int64, first int64, orderBy string, orderDirection OrderDirection, interceptors ...clientv2.RequestInterceptor) (*ListVotes, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli *http.Client, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) SnapshotClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type Query struct {
	Space         *Space          "json:\"space,omitempty\" graphql:\"space\""
	Spaces        []*Space        "json:\"spaces,omitempty\" graphql:\"spaces\""
	Ranking       *RankingObject  "json:\"ranking,omitempty\" graphql:\"ranking\""
	Proposal      *Proposal       "json:\"proposal,omitempty\" graphql:\"proposal\""
	Proposals     []*Proposal     "json:\"proposals,omitempty\" graphql:\"proposals\""
	Vote          *Vote           "json:\"vote,omitempty\" graphql:\"vote\""
	Votes         []*Vote         "json:\"votes,omitempty\" graphql:\"votes\""
	Aliases       []*Alias        "json:\"aliases,omitempty\" graphql:\"aliases\""
	Follows       []*Follow       "json:\"follows,omitempty\" graphql:\"follows\""
	Subscriptions []*Subscription "json:\"subscriptions,omitempty\" graphql:\"subscriptions\""
	Users         []*User         "json:\"users,omitempty\" graphql:\"users\""
	Statements    []*Statement    "json:\"statements,omitempty\" graphql:\"statements\""
	User          *User           "json:\"user,omitempty\" graphql:\"user\""
	Statement     *Statement      "json:\"statement,omitempty\" graphql:\"statement\""
	Skins         []*Item         "json:\"skins,omitempty\" graphql:\"skins\""
	Networks      []*Item         "json:\"networks,omitempty\" graphql:\"networks\""
	Validations   []*Item         "json:\"validations,omitempty\" graphql:\"validations\""
	Plugins       []*Item         "json:\"plugins,omitempty\" graphql:\"plugins\""
	Strategies    []*StrategyItem "json:\"strategies,omitempty\" graphql:\"strategies\""
	Strategy      *StrategyItem   "json:\"strategy,omitempty\" graphql:\"strategy\""
	Vp            *Vp             "json:\"vp,omitempty\" graphql:\"vp\""
	Messages      []*Message      "json:\"messages,omitempty\" graphql:\"messages\""
}
type NetworkFragment struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *NetworkFragment) GetID() string {
	if t == nil {
		t = &NetworkFragment{}
	}
	return t.ID
}

type ProposalFragment struct {
	ID            string                   "json:\"id\" graphql:\"id\""
	Ipfs          *string                  "json:\"ipfs,omitempty\" graphql:\"ipfs\""
	Author        string                   "json:\"author\" graphql:\"author\""
	Created       int64                    "json:\"created\" graphql:\"created\""
	Network       string                   "json:\"network\" graphql:\"network\""
	Symbol        string                   "json:\"symbol\" graphql:\"symbol\""
	Type          *string                  "json:\"type,omitempty\" graphql:\"type\""
	Strategies    []*StrategyFragment      "json:\"strategies\" graphql:\"strategies\""
	Validation    *ValidationFragment      "json:\"validation,omitempty\" graphql:\"validation\""
	Title         string                   "json:\"title\" graphql:\"title\""
	Body          *string                  "json:\"body,omitempty\" graphql:\"body\""
	Discussion    string                   "json:\"discussion\" graphql:\"discussion\""
	Choices       []*string                "json:\"choices\" graphql:\"choices\""
	Start         int64                    "json:\"start\" graphql:\"start\""
	End           int64                    "json:\"end\" graphql:\"end\""
	Quorum        float64                  "json:\"quorum\" graphql:\"quorum\""
	Privacy       *string                  "json:\"privacy,omitempty\" graphql:\"privacy\""
	Snapshot      *string                  "json:\"snapshot,omitempty\" graphql:\"snapshot\""
	State         *string                  "json:\"state,omitempty\" graphql:\"state\""
	Link          *string                  "json:\"link,omitempty\" graphql:\"link\""
	App           *string                  "json:\"app,omitempty\" graphql:\"app\""
	Scores        []*float64               "json:\"scores,omitempty\" graphql:\"scores\""
	ScoresState   *string                  "json:\"scores_state,omitempty\" graphql:\"scores_state\""
	ScoresTotal   *float64                 "json:\"scores_total,omitempty\" graphql:\"scores_total\""
	ScoresUpdated *int64                   "json:\"scores_updated,omitempty\" graphql:\"scores_updated\""
	Votes         *int64                   "json:\"votes,omitempty\" graphql:\"votes\""
	Flagged       *bool                    "json:\"flagged,omitempty\" graphql:\"flagged\""
	Space         *SpaceIdentifierFragment "json:\"space,omitempty\" graphql:\"space\""
}

func (t *ProposalFragment) GetID() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.ID
}
func (t *ProposalFragment) GetIpfs() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Ipfs
}
func (t *ProposalFragment) GetAuthor() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Author
}
func (t *ProposalFragment) GetCreated() int64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Created
}
func (t *ProposalFragment) GetNetwork() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Network
}
func (t *ProposalFragment) GetSymbol() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Symbol
}
func (t *ProposalFragment) GetType() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Type
}
func (t *ProposalFragment) GetStrategies() []*StrategyFragment {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Strategies
}
func (t *ProposalFragment) GetValidation() *ValidationFragment {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Validation
}
func (t *ProposalFragment) GetTitle() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Title
}
func (t *ProposalFragment) GetBody() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Body
}
func (t *ProposalFragment) GetDiscussion() string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Discussion
}
func (t *ProposalFragment) GetChoices() []*string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Choices
}
func (t *ProposalFragment) GetStart() int64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Start
}
func (t *ProposalFragment) GetEnd() int64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.End
}
func (t *ProposalFragment) GetQuorum() float64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Quorum
}
func (t *ProposalFragment) GetPrivacy() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Privacy
}
func (t *ProposalFragment) GetSnapshot() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Snapshot
}
func (t *ProposalFragment) GetState() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.State
}
func (t *ProposalFragment) GetLink() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Link
}
func (t *ProposalFragment) GetApp() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.App
}
func (t *ProposalFragment) GetScores() []*float64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Scores
}
func (t *ProposalFragment) GetScoresState() *string {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.ScoresState
}
func (t *ProposalFragment) GetScoresTotal() *float64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.ScoresTotal
}
func (t *ProposalFragment) GetScoresUpdated() *int64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.ScoresUpdated
}
func (t *ProposalFragment) GetVotes() *int64 {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Votes
}
func (t *ProposalFragment) GetFlagged() *bool {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Flagged
}
func (t *ProposalFragment) GetSpace() *SpaceIdentifierFragment {
	if t == nil {
		t = &ProposalFragment{}
	}
	return t.Space
}

type ProposalIdentifierFragment struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *ProposalIdentifierFragment) GetID() string {
	if t == nil {
		t = &ProposalIdentifierFragment{}
	}
	return t.ID
}

type SpaceFragment struct {
	ID              string               "json:\"id\" graphql:\"id\""
	Name            *string              "json:\"name,omitempty\" graphql:\"name\""
	Private         *bool                "json:\"private,omitempty\" graphql:\"private\""
	About           *string              "json:\"about,omitempty\" graphql:\"about\""
	Avatar          *string              "json:\"avatar,omitempty\" graphql:\"avatar\""
	Terms           *string              "json:\"terms,omitempty\" graphql:\"terms\""
	Location        *string              "json:\"location,omitempty\" graphql:\"location\""
	Website         *string              "json:\"website,omitempty\" graphql:\"website\""
	Twitter         *string              "json:\"twitter,omitempty\" graphql:\"twitter\""
	Github          *string              "json:\"github,omitempty\" graphql:\"github\""
	Coingecko       *string              "json:\"coingecko,omitempty\" graphql:\"coingecko\""
	Email           *string              "json:\"email,omitempty\" graphql:\"email\""
	Network         *string              "json:\"network,omitempty\" graphql:\"network\""
	Symbol          *string              "json:\"symbol,omitempty\" graphql:\"symbol\""
	Skin            *string              "json:\"skin,omitempty\" graphql:\"skin\""
	Domain          *string              "json:\"domain,omitempty\" graphql:\"domain\""
	ActiveProposals *int64               "json:\"activeProposals,omitempty\" graphql:\"activeProposals\""
	ProposalsCount  *int64               "json:\"proposalsCount,omitempty\" graphql:\"proposalsCount\""
	FollowersCount  *int64               "json:\"followersCount,omitempty\" graphql:\"followersCount\""
	VotesCount      *int64               "json:\"votesCount,omitempty\" graphql:\"votesCount\""
	Guidelines      *string              "json:\"guidelines,omitempty\" graphql:\"guidelines\""
	Template        *string              "json:\"template,omitempty\" graphql:\"template\""
	Verified        *bool                "json:\"verified,omitempty\" graphql:\"verified\""
	Flagged         *bool                "json:\"flagged,omitempty\" graphql:\"flagged\""
	Rank            *float64             "json:\"rank,omitempty\" graphql:\"rank\""
	Admins          []*string            "json:\"admins,omitempty\" graphql:\"admins\""
	Members         []*string            "json:\"members,omitempty\" graphql:\"members\""
	Moderators      []*string            "json:\"moderators,omitempty\" graphql:\"moderators\""
	Categories      []*string            "json:\"categories,omitempty\" graphql:\"categories\""
	Treasuries      []*TreasuryFragment  "json:\"treasuries,omitempty\" graphql:\"treasuries\""
	Voting          *SpaceVotingFragment "json:\"voting,omitempty\" graphql:\"voting\""
	Strategies      []*StrategyFragment  "json:\"strategies,omitempty\" graphql:\"strategies\""
}

func (t *SpaceFragment) GetID() string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.ID
}
func (t *SpaceFragment) GetName() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Name
}
func (t *SpaceFragment) GetPrivate() *bool {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Private
}
func (t *SpaceFragment) GetAbout() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.About
}
func (t *SpaceFragment) GetAvatar() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Avatar
}
func (t *SpaceFragment) GetTerms() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Terms
}
func (t *SpaceFragment) GetLocation() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Location
}
func (t *SpaceFragment) GetWebsite() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Website
}
func (t *SpaceFragment) GetTwitter() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Twitter
}
func (t *SpaceFragment) GetGithub() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Github
}
func (t *SpaceFragment) GetCoingecko() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Coingecko
}
func (t *SpaceFragment) GetEmail() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Email
}
func (t *SpaceFragment) GetNetwork() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Network
}
func (t *SpaceFragment) GetSymbol() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Symbol
}
func (t *SpaceFragment) GetSkin() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Skin
}
func (t *SpaceFragment) GetDomain() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Domain
}
func (t *SpaceFragment) GetActiveProposals() *int64 {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.ActiveProposals
}
func (t *SpaceFragment) GetProposalsCount() *int64 {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.ProposalsCount
}
func (t *SpaceFragment) GetFollowersCount() *int64 {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.FollowersCount
}
func (t *SpaceFragment) GetVotesCount() *int64 {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.VotesCount
}
func (t *SpaceFragment) GetGuidelines() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Guidelines
}
func (t *SpaceFragment) GetTemplate() *string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Template
}
func (t *SpaceFragment) GetVerified() *bool {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Verified
}
func (t *SpaceFragment) GetFlagged() *bool {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Flagged
}
func (t *SpaceFragment) GetRank() *float64 {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Rank
}
func (t *SpaceFragment) GetAdmins() []*string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Admins
}
func (t *SpaceFragment) GetMembers() []*string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Members
}
func (t *SpaceFragment) GetModerators() []*string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Moderators
}
func (t *SpaceFragment) GetCategories() []*string {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Categories
}
func (t *SpaceFragment) GetTreasuries() []*TreasuryFragment {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Treasuries
}
func (t *SpaceFragment) GetVoting() *SpaceVotingFragment {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Voting
}
func (t *SpaceFragment) GetStrategies() []*StrategyFragment {
	if t == nil {
		t = &SpaceFragment{}
	}
	return t.Strategies
}

type SpaceIdentifierFragment struct {
	ID string "json:\"id\" graphql:\"id\""
}

func (t *SpaceIdentifierFragment) GetID() string {
	if t == nil {
		t = &SpaceIdentifierFragment{}
	}
	return t.ID
}

type StrategyFragment struct {
	Name    string                 "json:\"name\" graphql:\"name\""
	Network *string                "json:\"network,omitempty\" graphql:\"network\""
	Params  map[string]interface{} "json:\"params,omitempty\" graphql:\"params\""
}

func (t *StrategyFragment) GetName() string {
	if t == nil {
		t = &StrategyFragment{}
	}
	return t.Name
}
func (t *StrategyFragment) GetNetwork() *string {
	if t == nil {
		t = &StrategyFragment{}
	}
	return t.Network
}
func (t *StrategyFragment) GetParams() map[string]interface{} {
	if t == nil {
		t = &StrategyFragment{}
	}
	return t.Params
}

type TreasuryFragment struct {
	Name    *string "json:\"name,omitempty\" graphql:\"name\""
	Address *string "json:\"address,omitempty\" graphql:\"address\""
	Network *string "json:\"network,omitempty\" graphql:\"network\""
}

func (t *TreasuryFragment) GetName() *string {
	if t == nil {
		t = &TreasuryFragment{}
	}
	return t.Name
}
func (t *TreasuryFragment) GetAddress() *string {
	if t == nil {
		t = &TreasuryFragment{}
	}
	return t.Address
}
func (t *TreasuryFragment) GetNetwork() *string {
	if t == nil {
		t = &TreasuryFragment{}
	}
	return t.Network
}

type ValidationFragment struct {
	Name   string                 "json:\"name\" graphql:\"name\""
	Params map[string]interface{} "json:\"params,omitempty\" graphql:\"params\""
}

func (t *ValidationFragment) GetName() string {
	if t == nil {
		t = &ValidationFragment{}
	}
	return t.Name
}
func (t *ValidationFragment) GetParams() map[string]interface{} {
	if t == nil {
		t = &ValidationFragment{}
	}
	return t.Params
}

type SpaceVotingFragment struct {
	Delay       *int64   "json:\"delay,omitempty\" graphql:\"delay\""
	Period      *int64   "json:\"period,omitempty\" graphql:\"period\""
	Type        *string  "json:\"type,omitempty\" graphql:\"type\""
	Quorum      *float64 "json:\"quorum,omitempty\" graphql:\"quorum\""
	Blind       *bool    "json:\"blind,omitempty\" graphql:\"blind\""
	HideAbstain *bool    "json:\"hideAbstain,omitempty\" graphql:\"hideAbstain\""
	Privacy     *string  "json:\"privacy,omitempty\" graphql:\"privacy\""
	Aliased     *bool    "json:\"aliased,omitempty\" graphql:\"aliased\""
}

func (t *SpaceVotingFragment) GetDelay() *int64 {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Delay
}
func (t *SpaceVotingFragment) GetPeriod() *int64 {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Period
}
func (t *SpaceVotingFragment) GetType() *string {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Type
}
func (t *SpaceVotingFragment) GetQuorum() *float64 {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Quorum
}
func (t *SpaceVotingFragment) GetBlind() *bool {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Blind
}
func (t *SpaceVotingFragment) GetHideAbstain() *bool {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.HideAbstain
}
func (t *SpaceVotingFragment) GetPrivacy() *string {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Privacy
}
func (t *SpaceVotingFragment) GetAliased() *bool {
	if t == nil {
		t = &SpaceVotingFragment{}
	}
	return t.Aliased
}

type VoteFragment struct {
	ID           string                      "json:\"id\" graphql:\"id\""
	Ipfs         *string                     "json:\"ipfs,omitempty\" graphql:\"ipfs\""
	Voter        string                      "json:\"voter\" graphql:\"voter\""
	Created      int64                       "json:\"created\" graphql:\"created\""
	Space        *SpaceIdentifierFragment    "json:\"space\" graphql:\"space\""
	Proposal     *ProposalIdentifierFragment "json:\"proposal,omitempty\" graphql:\"proposal\""
	Metadata     map[string]interface{}      "json:\"metadata,omitempty\" graphql:\"metadata\""
	Reason       *string                     "json:\"reason,omitempty\" graphql:\"reason\""
	App          *string                     "json:\"app,omitempty\" graphql:\"app\""
	Vp           *float64                    "json:\"vp,omitempty\" graphql:\"vp\""
	VpByStrategy []*float64                  "json:\"vp_by_strategy,omitempty\" graphql:\"vp_by_strategy\""
	VpState      *string                     "json:\"vp_state,omitempty\" graphql:\"vp_state\""
}

func (t *VoteFragment) GetID() string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.ID
}
func (t *VoteFragment) GetIpfs() *string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Ipfs
}
func (t *VoteFragment) GetVoter() string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Voter
}
func (t *VoteFragment) GetCreated() int64 {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Created
}
func (t *VoteFragment) GetSpace() *SpaceIdentifierFragment {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Space
}
func (t *VoteFragment) GetProposal() *ProposalIdentifierFragment {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Proposal
}
func (t *VoteFragment) GetMetadata() map[string]interface{} {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Metadata
}
func (t *VoteFragment) GetReason() *string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Reason
}
func (t *VoteFragment) GetApp() *string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.App
}
func (t *VoteFragment) GetVp() *float64 {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.Vp
}
func (t *VoteFragment) GetVpByStrategy() []*float64 {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.VpByStrategy
}
func (t *VoteFragment) GetVpState() *string {
	if t == nil {
		t = &VoteFragment{}
	}
	return t.VpState
}

type ListRanking_Ranking struct {
	Items []*SpaceIdentifierFragment "json:\"items,omitempty\" graphql:\"items\""
}

func (t *ListRanking_Ranking) GetItems() []*SpaceIdentifierFragment {
	if t == nil {
		t = &ListRanking_Ranking{}
	}
	return t.Items
}

type ListNetworks struct {
	Networks []*NetworkFragment "json:\"networks,omitempty\" graphql:\"networks\""
}

func (t *ListNetworks) GetNetworks() []*NetworkFragment {
	if t == nil {
		t = &ListNetworks{}
	}
	return t.Networks
}

type ListProposals struct {
	Proposals []*ProposalFragment "json:\"proposals,omitempty\" graphql:\"proposals\""
}

func (t *ListProposals) GetProposals() []*ProposalFragment {
	if t == nil {
		t = &ListProposals{}
	}
	return t.Proposals
}

type ProposalByID struct {
	Proposals []*ProposalFragment "json:\"proposals,omitempty\" graphql:\"proposals\""
}

func (t *ProposalByID) GetProposals() []*ProposalFragment {
	if t == nil {
		t = &ProposalByID{}
	}
	return t.Proposals
}

type ListRanking struct {
	Ranking *ListRanking_Ranking "json:\"ranking,omitempty\" graphql:\"ranking\""
}

func (t *ListRanking) GetRanking() *ListRanking_Ranking {
	if t == nil {
		t = &ListRanking{}
	}
	return t.Ranking
}

type ListSpaces struct {
	Spaces []*SpaceFragment "json:\"spaces,omitempty\" graphql:\"spaces\""
}

func (t *ListSpaces) GetSpaces() []*SpaceFragment {
	if t == nil {
		t = &ListSpaces{}
	}
	return t.Spaces
}

type SpaceByID struct {
	Spaces []*SpaceFragment "json:\"spaces,omitempty\" graphql:\"spaces\""
}

func (t *SpaceByID) GetSpaces() []*SpaceFragment {
	if t == nil {
		t = &SpaceByID{}
	}
	return t.Spaces
}

type ListVotes struct {
	Votes []*VoteFragment "json:\"votes,omitempty\" graphql:\"votes\""
}

func (t *ListVotes) GetVotes() []*VoteFragment {
	if t == nil {
		t = &ListVotes{}
	}
	return t.Votes
}

const ListNetworksDocument = `query ListNetworks {
	networks {
		... NetworkFragment
	}
}
fragment NetworkFragment on Item {
	id
}
`

func (c *Client) ListNetworks(ctx context.Context, interceptors ...clientv2.RequestInterceptor) (*ListNetworks, error) {
	vars := map[string]interface{}{}

	var res ListNetworks
	if err := c.Client.Post(ctx, "ListNetworks", ListNetworksDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListProposalsDocument = `query ListProposals ($skip: Int!, $first: Int!, $createdAfter: Int!, $spaces: [String], $orderBy: String!, $orderDirection: OrderDirection!) {
	proposals(skip: $skip, first: $first, where: {space_in:$spaces,created_gte:$createdAfter}, orderBy: $orderBy, orderDirection: $orderDirection) {
		... ProposalFragment
	}
}
fragment ProposalFragment on Proposal {
	id
	ipfs
	author
	created
	network
	symbol
	type
	strategies {
		... StrategyFragment
	}
	validation {
		... ValidationFragment
	}
	title
	body
	discussion
	choices
	start
	end
	quorum
	privacy
	snapshot
	state
	link
	app
	scores
	scores_state
	scores_total
	scores_updated
	votes
	flagged
	space {
		... SpaceIdentifierFragment
	}
}
fragment StrategyFragment on Strategy {
	name
	network
	params
}
fragment ValidationFragment on Validation {
	name
	params
}
fragment SpaceIdentifierFragment on Space {
	id
}
`

func (c *Client) ListProposals(ctx context.Context, skip int64, first int64, createdAfter int64, spaces []*string, orderBy string, orderDirection OrderDirection, interceptors ...clientv2.RequestInterceptor) (*ListProposals, error) {
	vars := map[string]interface{}{
		"skip":           skip,
		"first":          first,
		"createdAfter":   createdAfter,
		"spaces":         spaces,
		"orderBy":        orderBy,
		"orderDirection": orderDirection,
	}

	var res ListProposals
	if err := c.Client.Post(ctx, "ListProposals", ListProposalsDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ProposalByIDDocument = `query ProposalByID ($id: String!) {
	proposals(where: {id:$id}) {
		... ProposalFragment
	}
}
fragment ProposalFragment on Proposal {
	id
	ipfs
	author
	created
	network
	symbol
	type
	strategies {
		... StrategyFragment
	}
	validation {
		... ValidationFragment
	}
	title
	body
	discussion
	choices
	start
	end
	quorum
	privacy
	snapshot
	state
	link
	app
	scores
	scores_state
	scores_total
	scores_updated
	votes
	flagged
	space {
		... SpaceIdentifierFragment
	}
}
fragment StrategyFragment on Strategy {
	name
	network
	params
}
fragment ValidationFragment on Validation {
	name
	params
}
fragment SpaceIdentifierFragment on Space {
	id
}
`

func (c *Client) ProposalByID(ctx context.Context, id string, interceptors ...clientv2.RequestInterceptor) (*ProposalByID, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res ProposalByID
	if err := c.Client.Post(ctx, "ProposalByID", ProposalByIDDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListRankingDocument = `query ListRanking ($skip: Int!, $first: Int!, $category: String!, $network: String!) {
	ranking(skip: $skip, first: $first, where: {category:$category,network:$network}) {
		items {
			... SpaceIdentifierFragment
		}
	}
}
fragment SpaceIdentifierFragment on Space {
	id
}
`

func (c *Client) ListRanking(ctx context.Context, skip int64, first int64, category string, network string, interceptors ...clientv2.RequestInterceptor) (*ListRanking, error) {
	vars := map[string]interface{}{
		"skip":     skip,
		"first":    first,
		"category": category,
		"network":  network,
	}

	var res ListRanking
	if err := c.Client.Post(ctx, "ListRanking", ListRankingDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListSpacesDocument = `query ListSpaces ($skip: Int!, $first: Int!, $ids: [String]) {
	spaces(skip: $skip, first: $first, where: {id_in:$ids}) {
		... SpaceFragment
	}
}
fragment SpaceFragment on Space {
	id
	name
	private
	about
	avatar
	terms
	location
	website
	twitter
	github
	coingecko
	email
	network
	symbol
	skin
	domain
	activeProposals
	proposalsCount
	followersCount
	votesCount
	guidelines
	template
	verified
	flagged
	rank
	admins
	members
	moderators
	categories
	treasuries {
		... TreasuryFragment
	}
	voting {
		... SpaceVotingFragment
	}
	strategies {
		... StrategyFragment
	}
}
fragment TreasuryFragment on Treasury {
	name
	address
	network
}
fragment SpaceVotingFragment on SpaceVoting {
	delay
	period
	type
	quorum
	blind
	hideAbstain
	privacy
	aliased
}
fragment StrategyFragment on Strategy {
	name
	network
	params
}
`

func (c *Client) ListSpaces(ctx context.Context, skip int64, first int64, ids []*string, interceptors ...clientv2.RequestInterceptor) (*ListSpaces, error) {
	vars := map[string]interface{}{
		"skip":  skip,
		"first": first,
		"ids":   ids,
	}

	var res ListSpaces
	if err := c.Client.Post(ctx, "ListSpaces", ListSpacesDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const SpaceByIDDocument = `query SpaceByID ($id: String!) {
	spaces(where: {id:$id}) {
		... SpaceFragment
	}
}
fragment SpaceFragment on Space {
	id
	name
	private
	about
	avatar
	terms
	location
	website
	twitter
	github
	coingecko
	email
	network
	symbol
	skin
	domain
	activeProposals
	proposalsCount
	followersCount
	votesCount
	guidelines
	template
	verified
	flagged
	rank
	admins
	members
	moderators
	categories
	treasuries {
		... TreasuryFragment
	}
	voting {
		... SpaceVotingFragment
	}
	strategies {
		... StrategyFragment
	}
}
fragment TreasuryFragment on Treasury {
	name
	address
	network
}
fragment SpaceVotingFragment on SpaceVoting {
	delay
	period
	type
	quorum
	blind
	hideAbstain
	privacy
	aliased
}
fragment StrategyFragment on Strategy {
	name
	network
	params
}
`

func (c *Client) SpaceByID(ctx context.Context, id string, interceptors ...clientv2.RequestInterceptor) (*SpaceByID, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res SpaceByID
	if err := c.Client.Post(ctx, "SpaceByID", SpaceByIDDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListVotesDocument = `query ListVotes ($proposal: String!, $skip: Int!, $first: Int!, $orderBy: String!, $orderDirection: OrderDirection!) {
	votes(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection, where: {proposal:$proposal}) {
		... VoteFragment
	}
}
fragment VoteFragment on Vote {
	id
	ipfs
	voter
	created
	space {
		... SpaceIdentifierFragment
	}
	proposal {
		... ProposalIdentifierFragment
	}
	metadata
	reason
	app
	vp
	vp_by_strategy
	vp_state
}
fragment SpaceIdentifierFragment on Space {
	id
}
fragment ProposalIdentifierFragment on Proposal {
	id
}
`

func (c *Client) ListVotes(ctx context.Context, proposal string, skip int64, first int64, orderBy string, orderDirection OrderDirection, interceptors ...clientv2.RequestInterceptor) (*ListVotes, error) {
	vars := map[string]interface{}{
		"proposal":       proposal,
		"skip":           skip,
		"first":          first,
		"orderBy":        orderBy,
		"orderDirection": orderDirection,
	}

	var res ListVotes
	if err := c.Client.Post(ctx, "ListVotes", ListVotesDocument, &res, vars, interceptors...); err != nil {
		return nil, err
	}

	return &res, nil
}
