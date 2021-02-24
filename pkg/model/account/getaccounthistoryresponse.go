package account

import "github.com/shopspring/decimal"

type GetAccountHistoryResponse struct {
	Status string           `json:"status"`
	Data   []AccountHistory `json:"data"`
	NextId int64            `json:"next-id"`
}

type AccountHistory struct {
	AccountId    int64           `json:"account-id"`
	Currency     string          `json:"currency"`
	TransactAmt  decimal.Decimal `json:"transact-amt"`
	TransactType string          `json:"transact-type"`
	RecordId     int64           `json:"record-id"`
	AvailBalance decimal.Decimal `json:"avail-balance"`
	AcctBalance  decimal.Decimal `json:"acct-balance"`
	TransactTime int64           `json:"transact-time"`
}
