package order

import "github.com/shopspring/decimal"

type GetOrderResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         *struct {
		AccountId        int64           `json:"account-id"`
		Amount           decimal.Decimal `json:"amount"`
		Id               int64           `json:"id"`
		ClientOrderId    string          `json:"client-order-id"`
		Symbol           string          `json:"symbol"`
		Price            decimal.Decimal `json:"price"`
		CreatedAt        int64           `json:"created-at"`
		Type             string          `json:"type"`
		FilledAmount     decimal.Decimal `json:"field-amount"`
		FilledCashAmount decimal.Decimal `json:"field-cash-amount"`
		FilledFees       decimal.Decimal `json:"field-fees"`
		Source           string          `json:"source"`
		State            string          `json:"state"`
	}
}
