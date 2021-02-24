package order

import (
	"github.com/shopspring/decimal"
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type SubscribeOrderV2Response struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		EventType       string          `json:"eventType"`
		Symbol          string          `json:"symbol"`
		AccountId       int64           `json:"accountId"`
		OrderId         int64           `json:"orderId"`
		ClientOrderId   string          `json:"clientOrderId"`
		OrderSide       string          `json:"orderSide"`
		OrderPrice      decimal.Decimal `json:"orderPrice"`
		OrderSize       decimal.Decimal `json:"orderSize"`
		OrderValue      decimal.Decimal `json:"orderValue"`
		Type            string          `json:"type"`
		OrderStatus     string          `json:"orderStatus"`
		OrderCreateTime int64           `json:"orderCreateTime"`
		TradePrice      decimal.Decimal `json:"tradePrice"`
		TradeVolume     decimal.Decimal `json:"tradeVolume"`
		TradeId         int64           `json:"tradeId"`
		TradeTime       int64           `json:"tradeTime"`
		Aggressor       bool            `json:"aggressor"`
		RemainAmt       decimal.Decimal `json:"remainAmt"`
		LastActTime     int64           `json:"lastActTime"`
		ErrorCode       int             `json:"errCode"`
		ErrorMessage    string          `json:"errMessage"`
	}
}
