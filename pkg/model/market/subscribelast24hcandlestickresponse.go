package market

import (
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type SubscribeLast24hCandlestickResponse struct {
	base.WebSocketResponseBase
	Data *Candlestick
	Tick *Candlestick
}
