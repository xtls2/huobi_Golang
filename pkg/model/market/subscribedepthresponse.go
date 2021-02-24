package market

import (
	"github.com/xtls2/huobi_golang/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
