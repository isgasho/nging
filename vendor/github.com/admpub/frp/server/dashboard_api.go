// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/json"

	"github.com/admpub/frp/g"
	"github.com/admpub/frp/models/config"
	"github.com/admpub/frp/models/consts"
	"github.com/admpub/frp/utils/log"
	"github.com/admpub/frp/utils/version"
	"github.com/webx-top/echo"
)

type GeneralResponse struct {
	Code int
	Msg  string
}

type ServerInfoResp struct {
	Version           string `json:"version"`
	BindPort          int    `json:"bind_port"`
	BindUdpPort       int    `json:"bind_udp_port"`
	VhostHttpPort     int    `json:"vhost_http_port"`
	VhostHttpsPort    int    `json:"vhost_https_port"`
	KcpBindPort       int    `json:"kcp_bind_port"`
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

// api/serverinfo
func (svr *Service) ApiServerInfo(c echo.Context) error {
	res := GeneralResponse{Code: 200}

	log.Info("Http request: [%s]", c.Request().URL().Path())
	cfg := &g.GlbServerCfg.ServerCommonConf
	serverStats := svr.statsCollector.GetServer()
	svrResp := ServerInfoResp{
		Version:           version.Full(),
		BindPort:          cfg.BindPort,
		BindUdpPort:       cfg.BindUdpPort,
		VhostHttpPort:     cfg.VhostHttpPort,
		VhostHttpsPort:    cfg.VhostHttpsPort,
		KcpBindPort:       cfg.KcpBindPort,
		SubdomainHost:     cfg.SubDomainHost,
		MaxPoolCount:      cfg.MaxPoolCount,
		MaxPortsPerClient: cfg.MaxPortsPerClient,
		HeartBeatTimeout:  cfg.HeartBeatTimeout,

		TotalTrafficIn:  serverStats.TotalTrafficIn,
		TotalTrafficOut: serverStats.TotalTrafficOut,
		CurConns:        serverStats.CurConns,
		ClientCounts:    serverStats.ClientCounts,
		ProxyTypeCounts: serverStats.ProxyTypeCounts,
	}

	buf, _ := json.Marshal(&svrResp)
	res.Msg = string(buf)
	log.Info("Http response [%s]: code [%d]", c.Request().URL().Path(), res.Code)
	return c.String(res.Msg, res.Code)
}

type BaseOutConf struct {
	config.BaseProxyConf
}

type TcpOutConf struct {
	BaseOutConf
	RemotePort int `json:"remote_port"`
}

type UdpOutConf struct {
	BaseOutConf
	RemotePort int `json:"remote_port"`
}

type HttpOutConf struct {
	BaseOutConf
	config.DomainConf
	Locations         []string `json:"locations"`
	HostHeaderRewrite string   `json:"host_header_rewrite"`
}

type HttpsOutConf struct {
	BaseOutConf
	config.DomainConf
}

type StcpOutConf struct {
	BaseOutConf
}

type XtcpOutConf struct {
	BaseOutConf
}

func getConfByType(proxyType string) interface{} {
	switch proxyType {
	case consts.TcpProxy:
		return &TcpOutConf{}
	case consts.UdpProxy:
		return &UdpOutConf{}
	case consts.HttpProxy:
		return &HttpOutConf{}
	case consts.HttpsProxy:
		return &HttpsOutConf{}
	case consts.StcpProxy:
		return &StcpOutConf{}
	case consts.XtcpProxy:
		return &XtcpOutConf{}
	default:
		return nil
	}
}

// ProxyStatsInfo Get proxy info.
type ProxyStatsInfo struct {
	Name            string      `json:"name"`
	Conf            interface{} `json:"conf"`
	TodayTrafficIn  int64       `json:"today_traffic_in"`
	TodayTrafficOut int64       `json:"today_traffic_out"`
	CurConns        int64       `json:"cur_conns"`
	LastStartTime   string      `json:"last_start_time"`
	LastCloseTime   string      `json:"last_close_time"`
	Status          string      `json:"status"`
}

type GetProxyInfoResp struct {
	Proxies []*ProxyStatsInfo `json:"proxies"`
}

// api/proxy/:type
func (svr *Service) ApiProxyByType(c echo.Context) error {
	res := GeneralResponse{Code: 200}
	proxyType := c.Param("type")

	log.Info("Http request: [%s]", c.Request().URL().Path())

	proxyInfoResp := GetProxyInfoResp{}
	proxyInfoResp.Proxies = svr.getProxyStatsByType(proxyType)

	buf, _ := json.Marshal(&proxyInfoResp)
	res.Msg = string(buf)
	log.Info("Http response [%s]: code [%d]", c.Request().URL().Path(), res.Code)
	return c.String(res.Msg, res.Code)
}

func (svr *Service) getProxyStatsByType(proxyType string) (proxyInfos []*ProxyStatsInfo) {
	proxyStats := svr.statsCollector.GetProxiesByType(proxyType)
	proxyInfos = make([]*ProxyStatsInfo, 0, len(proxyStats))
	for _, ps := range proxyStats {
		proxyInfo := &ProxyStatsInfo{}
		if pxy, ok := svr.pxyManager.GetByName(ps.Name); ok {
			content, err := json.Marshal(pxy.GetConf())
			if err != nil {
				log.Warn("marshal proxy [%s] conf info error: %v", ps.Name, err)
				continue
			}
			proxyInfo.Conf = getConfByType(ps.Type)
			if err = json.Unmarshal(content, &proxyInfo.Conf); err != nil {
				log.Warn("unmarshal proxy [%s] conf info error: %v", ps.Name, err)
				continue
			}
			proxyInfo.Status = consts.Online
		} else {
			proxyInfo.Status = consts.Offline
		}
		proxyInfo.Name = ps.Name
		proxyInfo.TodayTrafficIn = ps.TodayTrafficIn
		proxyInfo.TodayTrafficOut = ps.TodayTrafficOut
		proxyInfo.CurConns = ps.CurConns
		proxyInfo.LastStartTime = ps.LastStartTime
		proxyInfo.LastCloseTime = ps.LastCloseTime
		proxyInfos = append(proxyInfos, proxyInfo)
	}
	return
}

// GetProxyStatsResp Get proxy info by name.
type GetProxyStatsResp struct {
	Name            string      `json:"name"`
	Conf            interface{} `json:"conf"`
	TodayTrafficIn  int64       `json:"today_traffic_in"`
	TodayTrafficOut int64       `json:"today_traffic_out"`
	CurConns        int64       `json:"cur_conns"`
	LastStartTime   string      `json:"last_start_time"`
	LastCloseTime   string      `json:"last_close_time"`
	Status          string      `json:"status"`
}

// api/proxy/:type/:name
func (svr *Service) ApiProxyByTypeAndName(c echo.Context) error {
	res := GeneralResponse{Code: 200}
	proxyType := c.Param("type")
	name := c.Param("name")

	log.Info("Http request: [%s]", c.Request().URL().Path())

	proxyStatsResp := GetProxyStatsResp{}
	proxyStatsResp, res.Code, res.Msg = svr.getProxyStatsByTypeAndName(proxyType, name)
	if res.Code != 200 {
		return c.String(res.Msg, res.Code)
	}

	buf, _ := json.Marshal(&proxyStatsResp)
	res.Msg = string(buf)
	log.Info("Http response [%s]: code [%d]", c.Request().URL().Path(), res.Code)
	return c.String(res.Msg, res.Code)
}

func (svr *Service) getProxyStatsByTypeAndName(proxyType string, proxyName string) (proxyInfo GetProxyStatsResp, code int, msg string) {
	proxyInfo.Name = proxyName
	ps := svr.statsCollector.GetProxiesByTypeAndName(proxyType, proxyName)
	if ps == nil {
		code = 404
		msg = "no proxy info found"
	} else {
		if pxy, ok := svr.pxyManager.GetByName(proxyName); ok {
			content, err := json.Marshal(pxy.GetConf())
			if err != nil {
				log.Warn("marshal proxy [%s] conf info error: %v", ps.Name, err)
				code = 400
				msg = "parse conf error"
				return
			}
			proxyInfo.Conf = getConfByType(ps.Type)
			if err = json.Unmarshal(content, &proxyInfo.Conf); err != nil {
				log.Warn("unmarshal proxy [%s] conf info error: %v", ps.Name, err)
				code = 400
				msg = "parse conf error"
				return
			}
			proxyInfo.Status = consts.Online
		} else {
			proxyInfo.Status = consts.Offline
		}
		proxyInfo.TodayTrafficIn = ps.TodayTrafficIn
		proxyInfo.TodayTrafficOut = ps.TodayTrafficOut
		proxyInfo.CurConns = ps.CurConns
		proxyInfo.LastStartTime = ps.LastStartTime
		proxyInfo.LastCloseTime = ps.LastCloseTime
		code = 200
	}

	return
}

// api/traffic/:name
type GetProxyTrafficResp struct {
	Name       string  `json:"name"`
	TrafficIn  []int64 `json:"traffic_in"`
	TrafficOut []int64 `json:"traffic_out"`
}

func (svr *Service) ApiProxyTraffic(c echo.Context) error {
	res := GeneralResponse{Code: 200}
	name := c.Param("name")
	log.Info("Http request: [%s]", c.Request().URL().Path())

	trafficResp := GetProxyTrafficResp{}
	trafficResp.Name = name
	proxyTrafficInfo := svr.statsCollector.GetProxyTraffic(name)

	if proxyTrafficInfo == nil {
		res.Code = 404
		res.Msg = "no proxy info found"
		return c.String(res.Msg, res.Code)
	} else {
		trafficResp.TrafficIn = proxyTrafficInfo.TrafficIn
		trafficResp.TrafficOut = proxyTrafficInfo.TrafficOut
	}

	buf, _ := json.Marshal(&trafficResp)
	res.Msg = string(buf)
	log.Info("Http response [%s]: code [%d]", c.Request().URL().Path(), res.Code)
	return c.String(res.Msg, res.Code)
}
