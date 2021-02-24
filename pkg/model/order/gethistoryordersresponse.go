package order

import "github.com/shopspring/decimal"

type HistoryOrder struct {
	Id               int64           `json:"id"`
	ClientOrderId    string          `json:"client-order-id"`
	AccountId        int64           `json:"account-id"`
	UserId           int             `json:"user-id"`
	Amount           decimal.Decimal `json:"amount"`
	Symbol           string          `json:"symbol"`
	Price            decimal.Decimal `json:"price"`
	CreatedAt        int64           `json:"created-at"`
	CanceledAt       int64           `json:"canceled-at"`
	FinishedAt       int64           `json:"finished-at"`
	Type             string          `json:"type"`
	FilledAmount     decimal.Decimal `json:"field-amount"`
	FilledCashAmount decimal.Decimal `json:"field-cash-amount"`
	FilledFees       decimal.Decimal `json:"field-fees"`
	Source           string          `json:"source"`
	State            string          `json:"state"`
	Exchange         string          `json:"exchange"`
	Batch            string          `json:"batch"`
	StopPrice        decimal.Decimal `json:"stop-price"`
	Operator         string          `json:"operator"`
}

type GetHistoryOrdersResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []HistoryOrder
}
