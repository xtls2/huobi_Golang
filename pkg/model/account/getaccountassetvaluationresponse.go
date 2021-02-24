package account

import "github.com/shopspring/decimal"

type GetAccountAssetValuationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		Balance   decimal.Decimal `json:"balance"`
		Timestamp int64           `json:"timestamp"`
	}
}
