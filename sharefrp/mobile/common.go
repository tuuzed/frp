package sharedfrp

import "C"
import (
	"github.com/fatedier/frp/pkg/util/version"
	"github.com/fatedier/frp/sharefrp/pkg/logx"
)

const (
	STATUS_NONE = iota
	STATUS_RUNNING
	STATUS_STOP
)
const (
	LOG_LEVEL_TRACE = logx.Trace
	LOG_LEVEL_DEBUG = logx.Debug
	LOG_LEVEL_INFO  = logx.Info
	LOG_LEVEL_WARN  = logx.Warn
	LOG_LEVEL_ERROR = logx.Error
)

type LogListener interface {
	Log(level int, msg string)
}

func ListenLog(l LogListener) {
	logx.AddLogAppender(l)
}

func GetFrpVersion() string {
	return version.Full()
}

func GetSharedFrpVersion() string {
	return GetFrpVersion() + "-20200415"
}
