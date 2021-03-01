// shared frp
package server

import (
	"errors"
	"github.com/fatedier/frp/pkg/metrics/mem"
	"github.com/fatedier/frp/pkg/util/version"
)

type ServerInfo struct {
	Version           string `json:"version"`
	BindPort          int    `json:"bind_port"`
	BindUDPPort       int    `json:"bind_udp_port"`
	VhostHTTPPort     int    `json:"vhost_http_port"`
	VhostHTTPSPort    int    `json:"vhost_https_port"`
	KCPBindPort       int    `json:"kcp_bind_port"`
	SubdomainHost     string `json:"subdomain_host"`
	MaxPoolCount      int64  `json:"max_pool_count"`
	MaxPortsPerClient int64  `json:"max_ports_per_client"`
	HeartBeatTimeout  int64  `json:"heart_beat_timeout"`

	TotalTrafficIn  int64            `json:"total_traffic_in"`
	TotalTrafficOut int64            `json:"total_traffic_out"`
	CurConns        int64            `json:"cur_conns"`
	ClientCounts    int64            `json:"client_counts"`
	ProxyTypeCounts map[string]int64 `json:"proxy_type_count"`
}

func (svr *Service) GetServerInfo() (serverInfo ServerInfo) {
	serverStats := mem.StatsCollector.GetServer()
	serverInfo = ServerInfo{
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

func (svr *Service) GetProxyStatsByTypeAndName(proxyType string, proxyName string) (proxyInfo GetProxyStatsResp, err error) {
	proxyInfo, code, msg := svr.getProxyStatsByTypeAndName(proxyType, proxyName)
	if code != 200 {
		err = errors.New(msg)
		return
	}
	return
}
