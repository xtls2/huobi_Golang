package order

import (
	"github.com/shopspring/decimal"
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type SubscribeTradeClearResponse struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		EventType       string          `json:"eventType"`
		Symbol          string          `json:"symbol"`
		OrderId         int64           `json:"orderId"`
		TradePrice      decimal.Decimal `json:"tradePrice"`
		TradeVolume     decimal.Decimal `json:"tradeVolume"`
		OrderSide       string          `json:"orderSide"`
		OrderType       string          `json:"orderType"`
		Aggressor       bool            `json:"aggressor"`
		TradeId         int64           `json:"tradeId"`
		TradeTime       int64           `json:"tradeTime"`
		TransactFee     decimal.Decimal `json:"transactFee"`
		FeeCurrency     decimal.Decimal `json:"feeCurrency"`
		FeeDeduct       string          `json:"feeDeduct"`
		FeeDeductType   string          `json:"feeDeductType"`
		AccountId       int64           `json:"accountId"`
		Source          string          `json:"source"`
		OrderPrice      decimal.Decimal `json:"orderPrice"`
		OrderSize       decimal.Decimal `json:"orderSize"`
		OrderValue      decimal.Decimal `json:"orderValue"`
		ClientOrderId   string          `json:"clientOrderId"`
		StopPrice       decimal.Decimal `json:"stopPrice"`
		Operator        string          `json:"operator"`
		OrderCreateTime int64           `json:"orderCreateTime"`
		OrderStatus     string          `json:"orderStatus"`
	}
}
