package client

import (
	"github.com/fatedier/frp/client/proxy"
)

func (svr *Service) GetAllProxyStatus() []*proxy.WorkingStatus {
	return svr.ctl.pm.GetAllProxyStatus()
}
