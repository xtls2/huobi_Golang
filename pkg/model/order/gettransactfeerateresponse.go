package order

import "github.com/shopspring/decimal"

type TransactFeeRate struct {
	Symbol          string          `json:"symbol"`
	MakerFeeRate    decimal.Decimal `json:"makerFeeRate"`
	TakerFeeRate    decimal.Decimal `json:"takerFeeRate"`
	ActualMakerRate decimal.Decimal `json:"actualMakerRate"`
	ActualTakerRate decimal.Decimal `json:"actualTakerRate"`
}

type GetTransactFeeRateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []TransactFeeRate
}
