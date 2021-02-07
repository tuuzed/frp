package sharedfrp

import (
	"context"
	"errors"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/sharefrp/pkg/assets"
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	kcpDoneCh                                 = make(chan struct{})
	clientServices                            = make(map[string]*client.Service)
	clientStatuses                            = make(map[string]int)
	clientStatusListener ClientStatusListener = nil
)

type ClientStatusListener interface {
	Status(name string, status int)
}

func SetClientStatusListener(l ClientStatusListener) {
	clientStatusListener = l
}

func LoadClientAssets() {
	assets.LoadClientAssets()
}

func GetClientStatus(name string) int {
	sc := clientStatuses[name]
	if sc == 0 {
		return STATUS_NONE
	}
	return sc
}

func RunClient(name string, iniContent, iniPath string) (err error) {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	clientStatuses[name] = STATUS_NONE
	if clientStatusListener != nil {
		clientStatusListener.Status(name, STATUS_NONE)
	}

	cfg, err := config.UnmarshalClientConfFromIni(iniContent)
	if err != nil {
		return
	}
	pxyCfgs, visitorCfgs, err := config.LoadAllConfFromIni(cfg.User, iniContent, cfg.Start)
	if err != nil {
		return err
	}
	err = startService(name, cfg, pxyCfgs, visitorCfgs, iniPath)
	return
}

func CloseClient(name string) (err error) {
	svr := clientServices[name]
	if svr == nil {
		err = errors.New("Not name ")
		return
	}
	svr.Close()
	delete(clientServices, name)
	return
}

func ReloadClientConf(name, iniContent string) (err error) {
	svr := clientServices[name]
	if svr == nil {
		err = errors.New("Not name ")
		return
	}
	cfg, err := config.UnmarshalClientConfFromIni(iniContent)
	if err != nil {
		return
	}
	pxyCfgs, visitorCfgs, err := config.LoadAllConfFromIni(cfg.User, iniContent, cfg.Start)
	if err != nil {
		return err
	}
	err = svr.ReloadConf(pxyCfgs, visitorCfgs)
	return
}

func handleSignal(svr *client.Service) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.Close()
	time.Sleep(250 * time.Millisecond)
	close(kcpDoneCh)
}

func startService(name string, cfg config.ClientCommonConf, pxyCfgs map[string]config.ProxyConf, visitorCfgs map[string]config.VisitorConf, cfgFile string) (err error) {
	defer func() {
		clientStatuses[name] = STATUS_STOP
		if clientStatusListener != nil {
			clientStatusListener.Status(name, STATUS_STOP)
		}
		delete(clientServices, name)
	}()
	log.InitLog(cfg.LogWay, cfg.LogFile, cfg.LogLevel,
		cfg.LogMaxDays, cfg.DisableLogColor)

	if cfg.DNSServer != "" {
		s := cfg.DNSServer
		if !strings.Contains(s, ":") {
			s += ":53"
		}
		// Change default dns server for frpc
		net.DefaultResolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("udp", s)
			},
		}
	}
	svr, errRet := client.NewService(cfg, pxyCfgs, visitorCfgs, cfgFile)
	if errRet != nil {
		err = errRet
		return
	}
	clientServices[name] = svr
	clientStatuses[name] = STATUS_RUNNING
	if clientStatusListener != nil {
		clientStatusListener.Status(name, STATUS_RUNNING)
	}
	// Capture the exit signal if we use kcp.
	if cfg.Protocol == "kcp" {
		go handleSignal(svr)
	}
	err = svr.Run()
	if cfg.Protocol == "kcp" {
		<-kcpDoneCh
	}
	return
}
