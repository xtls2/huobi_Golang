package algoorder

import "github.com/shopspring/decimal"

type GetHistoryOrdersResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		AccountId       int             `json:"accountId"`
		Source          string          `json:"source"`
		ClientOrderId   string          `json:"clientOrderId"`
		Symbol          string          `json:"symbol"`
		OrderPrice      decimal.Decimal `json:"orderPrice"`
		OrderSize       decimal.Decimal `json:"orderSize"`
		OrderValue      decimal.Decimal `json:"orderValue"`
		OrderSide       string          `json:"orderSide"`
		TimeInForce     string          `json:"timeInForce"`
		OrderType       string          `json:"orderType"`
		StopPrice       decimal.Decimal `json:"stopPrice"`
		TrailingRate    decimal.Decimal `json:"trailingRate"`
		OrderOrigTime   int64           `json:"orderOrigTime"`
		LastActTime     int64           `json:"lastActTime"`
		OrderCreateTime int64           `json:"orderCreateTime"`
		OrderStatus     string          `json:"orderStatus"`
		ErrCode         int             `json:"errCode"`
		ErrMessage      string          `json:"errMessage"`
	}
	NextId int64 `json:"nextId"`
}
