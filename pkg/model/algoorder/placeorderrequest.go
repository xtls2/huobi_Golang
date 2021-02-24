package algoorder

import "github.com/shopspring/decimal"

type PlaceOrderRequest struct {
	AccountId     int64           `json:"accountId"`
	Symbol        string          `json:"symbol"`
	OrderPrice    decimal.Decimal `json:"orderPrice"`
	OrderSide     string          `json:"orderSide"`
	OrderSize     decimal.Decimal `json:"orderSize"`
	OrderValue    decimal.Decimal `json:"orderValue"`
	TimeInForce   string          `json:"timeInForce"`
	OrderType     string          `json:"orderType"`
	ClientOrderId string          `json:"clientOrderId"`
	StopPrice     decimal.Decimal `json:"stopPrice"`
	TrailingRate  decimal.Decimal `json:"trailingRate"`
}
