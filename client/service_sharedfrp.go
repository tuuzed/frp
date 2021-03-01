// shared frp
package client

import (
	"errors"
	"github.com/fatedier/frp/client/proxy"
)

func (svr *Service) GetAllProxyStatus() (status []*proxy.WorkingStatus, err error) {
	ctl := svr.GetController()
	if ctl != nil {
		status = ctl.pm.GetAllProxyStatus()
		return
	}
	err = errors.New("not login")
	return
}
