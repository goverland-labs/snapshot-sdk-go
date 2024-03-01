package snapshot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/goverland-labs/snapshot-sdk-go/client"
)

type ValidationParams struct {
	Validation string                 `json:"validation"`
	Author     string                 `json:"author"`
	Space      string                 `json:"space"`
	Network    string                 `json:"network"`
	Snapshot   any                    `json:"snapshot"` // int or 'latest'
	Params     map[string]interface{} `json:"params"`
}

type ValidationResponse struct {
	Result bool `json:"result"`
}

type GetVotingPowerParams struct {
	Voter    string `json:"voter"`
	Space    string `json:"space"`
	Proposal string `json:"proposal"`
}

type StrategyFragment struct {
	Name    string                 `json:"name"`
	Network *string                `json:"network,omitempty"`
	Params  map[string]interface{} `json:"params"`
}

type VoteParams struct {
	Address string `json:"address"`
	Sig     string `json:"sig"`
	Data    any    `json:"data"`
}

type VoteResult struct {
	ID      string  `json:"id"`
	IPFS    string  `json:"ipfs"`
	Relayer Relayer `json:"relayer"`
}

type Relayer struct {
	Address string `json:"address"`
	Receipt string `json:"receipt"`
}

type jsonRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
}

func (s *SDK) Validate(_ context.Context, params ValidationParams) (ValidationResponse, error) {
	resp, err := s.doJSONRPCRequest("validate", params)
	if err != nil {
		return ValidationResponse{}, err
	}

	var result ValidationResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return ValidationResponse{}, err
	}

	return result, nil
}

func (s *SDK) GetVotingPower(ctx context.Context, params GetVotingPowerParams) (*client.VotingPowerFragment, error) {
	vp, err := wrapError(s.client.GetVotingPower(ctx, params.Voter, params.Space, params.Proposal))
	if err != nil {
		return nil, err
	}

	return vp.GetVp(), nil
}

func (s *SDK) Vote(_ context.Context, params VoteParams) (VoteResult, error) {
	resp, err := s.doSimpleRequest(params)
	if err != nil {
		return VoteResult{}, err
	}

	var result VoteResult
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return VoteResult{}, err
	}

	return result, nil
}

// TODO: correct handle 400
func (s *SDK) doJSONRPCRequest(method string, params any) ([]byte, error) {
	// Create a HTTP post request
	rpcReq, err := json.Marshal(jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	})
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", s.scoreURL, bytes.NewBuffer(rpcReq))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	res, err := s.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("json rpc request failed: %s", string(content))
	}

	return content, nil
}

func (s *SDK) doSimpleRequest(params any) ([]byte, error) {
	reqParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", s.seqURL, bytes.NewBuffer(reqParams))
	if err != nil {
		return nil, err
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	res, err := s.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("json rpc request failed: %s", string(content))
	}

	return content, nil
}
