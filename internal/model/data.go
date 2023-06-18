package model

type DataHeadOutput struct {
	TodayOrderCount int `json:"today_order_count" dc:"今日订单总数"`
	DAU             int `json:"dau" dc:"今日日活"`
	ConversionRate  int `json:"conversion_rate" dc:"转化率"`
}
