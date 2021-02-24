package market

import (
	"github.com/shopspring/decimal"
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type BestBidOffer struct {
	QuoteTime int64           `json:"quoteTime"`
	Symbol    string          `json:"symbol"`
	Bid       decimal.Decimal `json:"bid"`
	BidSize   decimal.Decimal `json:"bidSize"`
	Ask       decimal.Decimal `json:"ask"`
	AskSize   decimal.Decimal `json:"askSize"`
}
type SubscribeBestBidOfferResponse struct {
	base.WebSocketResponseBase
	Tick *BestBidOffer
}
