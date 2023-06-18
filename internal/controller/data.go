package controller

import (
	"context"
	"rime-shop-gf/api/backend"
	"rime-shop-gf/internal/service"
)

type cData struct{}

var Data = cData{}

func (c *cData) DataHead(ctx context.Context, req *backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	out, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DataHeadRes{
		TodayOrderCount: out.TodayOrderCount,
		DAU:             out.DAU,
		ConversionRate:  out.ConversionRate,
	}, nil
}
