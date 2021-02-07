package sharedfrp

import (
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
	"github.com/fatedier/frp/sharefrp/pkg/assets"
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"time"
)

var (
	serverStatus                             = STATUS_NONE
	serverStateListener ServerStatusListener = nil
)

type ServerStatusListener interface {
	Status(status int)
}

func SetServerStatusListener(l ServerStatusListener) {
	serverStateListener = l
}

func LoadServerAssets() {
	assets.LoadServerAssets()
}

func GetServerStatus() int {
	return serverStatus
}

func RunServer(iniContent string) (err error) {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	serverStatus = STATUS_NONE
	if serverStateListener != nil {
		serverStateListener.Status(STATUS_NONE)
	}

	cfg, err := config.UnmarshalServerConfFromIni(iniContent)
	if err != nil {
		return
	}
	err = runServer(cfg)
	return
}

func runServer(cfg config.ServerCommonConf) (err error) {
	defer func() {
		serverStatus = STATUS_STOP
		if serverStateListener != nil {
			serverStateListener.Status(STATUS_STOP)
		}
	}()
	log.InitLog(cfg.LogWay, cfg.LogFile, cfg.LogLevel, cfg.LogMaxDays, cfg.DisableLogColor)
	svr, err := server.NewService(cfg)
	if err != nil {
		return err
	}
	log.Info("start frps success")
	serverStatus = STATUS_RUNNING
	if serverStateListener != nil {
		serverStateListener.Status(STATUS_RUNNING)
	}
	svr.Run()
	return
}
