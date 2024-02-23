package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type BaseHoldDataService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *BaseHoldDataService) Symbol(symbol string) *BaseHoldDataService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *BaseHoldDataService) Period(period string) *BaseHoldDataService {
	s.period = period
	return s
}

// Limit set limit
func (s *BaseHoldDataService) Limit(limit int) *BaseHoldDataService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BaseHoldDataService) StartTime(startTime int64) *BaseHoldDataService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BaseHoldDataService) EndTime(endTime int64) *BaseHoldDataService {
	s.endTime = &endTime
	return s
}

func (s *BaseHoldDataService) Request(uri string, ctx context.Context, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: uri,
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	return s.c.callAPI(ctx, r, opts...)
}

func handleLongShortRatioData(data []byte, err error) ([]*LongShortRatio, error) {
	if err != nil {
		return []*LongShortRatio{}, err
	}

	res := make([]*LongShortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*LongShortRatio{}, err
	}

	return res, nil

}

// Top Trader Long/Short Ratio (Accounts)
// GET /futures/data/topLongShortAccountRatio
type TopTraderLongShortAccountRatioService struct {
	BaseHoldDataService
}

func (t *TopTraderLongShortAccountRatioService) Do(ctx context.Context, opts ...RequestOption) ([]*LongShortRatio, error) {
	const uri = "/futures/data/topLongShortAccountRatio"
	data, _, err := t.Request(uri, ctx, opts...)
	return handleLongShortRatioData(data, err)
}

// Top Trader Long/Short Ratio (Positions)
// GET /futures/data/topLongShortPositionRatio
type TopTraderLongShortPositionsRatioService struct {
	BaseHoldDataService
}

func (t *TopTraderLongShortPositionsRatioService) Do(ctx context.Context, opts ...RequestOption) ([]*LongShortRatio, error) {
	const uri = "/futures/data/topLongShortPositionRatio"
	data, _, err := t.Request(uri, ctx, opts...)
	return handleLongShortRatioData(data, err)
}
