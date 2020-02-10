package assets

import (
	c "github.com/fatedier/frp/assets/frpc/statik"
	s "github.com/fatedier/frp/assets/frps/statik"
)

func LoadServerAssets() {
	s.Load()
}

func LoadClientAssets() {
	c.Load()
}
