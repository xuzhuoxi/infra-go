package wsx

import "github.com/xuzhuoxi/infra-go/bytex"

// WsDataBlockHandler is the default data block handler for WebSocket pack framing.
// WsDataBlockHandler 是 WebSocket 封包默认的数据块处理器。
var WsDataBlockHandler = bytex.NewDefaultDataBlockHandler()
