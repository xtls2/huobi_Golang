package order

import "github.com/shopspring/decimal"

type MatchResults struct {
	Id                int64           `json:"id"`
	OrderId           int64           `json:"order-id"`
	MatchId           int64           `json:"match-id"`
	TradeId           int64           `json:"trade-id"`
	Symbol            string          `json:"symbol"`
	Price             decimal.Decimal `json:"price"`
	CreatedAt         int64           `json:"created-at"`
	Type              string          `json:"type"`
	FilledAmount      decimal.Decimal `json:"filled-amount"`
	FilledFees        decimal.Decimal `json:"filled-fees"`
	FeeCurrency       string          `json:"fee-currency"`
	Source            string          `json:"source"`
	Role              string          `json:"role"`
	FilledPoints      decimal.Decimal `json:"filled-points"`
	FeeDeductCurrency string          `json:"fee-deduct-currency"`
	FeeDeductState    string          `json:"fee-deduct-state"`
}

type GetMatchResultsResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Data         []MatchResults
}
