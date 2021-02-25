package algoorder

import "github.com/shopspring/decimal"

type GetOpenOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		AccountId     int             `json:"accountId"`
		Source        string          `json:"source"`
		ClientOrderId string          `json:"clientOrderId"`
		Symbol        string          `json:"symbol"`
		OrderPrice    decimal.Decimal `json:"orderPrice"`
		OrderSize     decimal.Decimal `json:"orderSize"`
		OrderValue    decimal.Decimal `json:"orderValue"`
		OrderSide     string          `json:"orderSide"`
		TimeInForce   string          `json:"timeInForce"`
		OrderType     string          `json:"orderType"`
		StopPrice     string          `json:"stopPrice"`
		TrailingRate  decimal.Decimal `json:"trailingRate"`
		OrderOrigTime int64           `json:"orderOrigTime"`
		LastActTime   int64           `json:"lastActTime"`
		OrderStatus   string          `json:"orderStatus"`
	}
	NextId int64 `json:"nextId"`
}
