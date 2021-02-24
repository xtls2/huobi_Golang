package account

import (
	"github.com/shopspring/decimal"
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type SubscribeAccountV2Response struct {
	base.WebSocketV2ResponseBase
	Data *struct {
		Currency    string          `json:"currency"`
		AccountId   int64           `json:"accountId"`
		Balance     decimal.Decimal `json:"balance"`
		Available   decimal.Decimal `json:"available"`
		ChangeType  string          `json:"changeType"`
		AccountType string          `json:"accountType"`
		ChangeTime  int64           `json:"changeTime"`
	}
}
