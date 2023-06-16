package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"rime-shop-gf/internal/dao"
	"rime-shop-gf/internal/model"
	"rime-shop-gf/internal/service"
	"rime-shop-gf/utility"
)

type sData struct {
}

// 注册服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {
	return &model.DataHeadOutput{
		TodayOrderCount: todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

// 查询今天的订单总数
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).WhereBetween(
		dao.OrderInfo.Columns().CreatedAt,
		gtime.Now().StartOfDay(),
		gtime.Now().EndOfDay(),
	).Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		// 返回 -1 说明查询有误
		return -1
	}
	return
}
