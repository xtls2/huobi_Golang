package order

import "github.com/shopspring/decimal"

type PlaceOrderRequest struct {
	AccountId     string          `json:"account-id"`
	Symbol        string          `json:"symbol"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount"`
	Price         decimal.Decimal `json:"price,omitempty"`
	Source        string          `json:"source,omitempty"`
	ClientOrderId string          `json:"client-order-id,omitempty"`
	StopPrice     decimal.Decimal `json:"stop-price,omitempty"`
	Operator      string          `json:"operator,omitempty"`
}
