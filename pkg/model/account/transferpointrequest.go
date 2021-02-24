package account

import "github.com/shopspring/decimal"

type TransferPointRequest struct {
	FromUid string          `json:"fromUid"`
	ToUid   string          `json:"toUid"`
	GroupId int64           `json:"groupId"`
	Amount  decimal.Decimal `json:"amount"`
}
