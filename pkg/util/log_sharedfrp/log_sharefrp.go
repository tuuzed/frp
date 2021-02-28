package log_sharedfrp

import (
	"github.com/fatedier/beego/logs"
	"sync"
)

var (
	appenderList = make([]Appender, 0)
	lock         = &sync.Mutex{}
)

const (
	Trace = logs.LevelTrace
	Debug = logs.LevelDebug
	Info  = logs.LevelInfo
	Warn  = logs.LevelWarn
	Error = logs.LevelError
)

type Appender interface {
	Log(level int, msg string)
}

func AddLogAppender(appender Appender) {
	lock.Lock()
	defer lock.Unlock()
	appenderList = append(appenderList, appender)
}

func Log(level int, msg string) {
	lock.Lock()
	defer lock.Unlock()
	for _, l := range appenderList {
		l.Log(level, msg)
	}
}
