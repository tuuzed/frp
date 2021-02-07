package assets

import (
	"github.com/fatedier/frp/sharefrp/pkg/assets/frpc"
	"github.com/fatedier/frp/sharefrp/pkg/assets/frps"
)

func LoadServerAssets() {
	frps.Load()
}

func LoadClientAssets() {
	frpc.Load()
}
