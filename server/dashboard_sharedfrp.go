// shared frp
package server

import (
	"errors"
	"github.com/fatedier/frp/pkg/metrics/mem"
	"github.com/fatedier/frp/pkg/util/version"
)

type ServerInfo serverInfoResp

func (svr *Service) GetServerInfo() (serverInfo *ServerInfo) {
	serverStats := mem.StatsCollector.GetServer()
	serverInfo = &ServerInfo{
		Version:           version.Full(),
		BindPort:          svr.cfg.BindPort,
		BindUDPPort:       svr.cfg.BindUDPPort,
		VhostHTTPPort:     svr.cfg.VhostHTTPPort,
		VhostHTTPSPort:    svr.cfg.VhostHTTPSPort,
		KCPBindPort:       svr.cfg.KCPBindPort,
		SubdomainHost:     svr.cfg.SubDomainHost,
		MaxPoolCount:      svr.cfg.MaxPoolCount,
		MaxPortsPerClient: svr.cfg.MaxPortsPerClient,
		HeartBeatTimeout:  svr.cfg.HeartbeatTimeout,

		TotalTrafficIn:  serverStats.TotalTrafficIn,
		TotalTrafficOut: serverStats.TotalTrafficOut,
		CurConns:        serverStats.CurConns,
		ClientCounts:    serverStats.ClientCounts,
		ProxyTypeCounts: serverStats.ProxyTypeCounts,
	}
	return
}

func (svr *Service) GetProxyStatsByType(proxyType string) (proxyInfos []*ProxyStatsInfo) {
	return svr.getProxyStatsByType(proxyType)
}

func (svr *Service) GetProxyStatsByTypeAndName(proxyType, proxyName string) (proxyInfo GetProxyStatsResp, err error) {
	proxyInfo, code, msg := svr.getProxyStatsByTypeAndName(proxyType, proxyName)
	if code != 200 {
		err = errors.New(msg)
		return
	}
	return
}
