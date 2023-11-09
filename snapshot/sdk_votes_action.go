package snapshot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Address    string             `json:"address"`
	Network    string             `json:"network"`
	Strategies []StrategyFragment `json:"strategies"`
	Snapshot   any                `json:"snapshot"` // int or 'latest'
	Space      string             `json:"space"`
	Delegation bool               `json:"delegation"`
}

type StrategyFragment struct {
	Name    string                 `json:"name"`
	Network *string                `json:"network,omitempty"`
	Params  map[string]interface{} `json:"params"`
}

type GetVotingPowerResponse struct {
	Result VotingPower `json:"result"`
}

type VotingPower struct {
	VP           float64   `json:"vp"`
	VPByStrategy []float64 `json:"vp_by_strategy"`
	VpState      string    `json:"vp_state"`
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

func (s *SDK) GetVotingPower(_ context.Context, params GetVotingPowerParams) (GetVotingPowerResponse, error) {
	resp, err := s.doJSONRPCRequest("get_vp", params)
	if err != nil {
		return GetVotingPowerResponse{}, err
	}

	var result GetVotingPowerResponse
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return GetVotingPowerResponse{}, err
	}

	return result, nil
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
