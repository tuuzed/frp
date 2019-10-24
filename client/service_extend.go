package client

func (svr *Service) IsExited() bool {
	return svr.exit != 0
}
