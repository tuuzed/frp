package log

import (
	"sync"
)

var (
	appenderList = make([]Appender, 0)
	lock         = &sync.Mutex{}
)

const (
	TraceLevel = 8
	DebugLevel = 7
	InfoLevel  = 6
	WarnLevel  = 4
	ErrorLevel = 3
)

type Appender interface {
	Log(level int, msg string)
}

func AddLogAppender(appender Appender) {
	lock.Lock()
	appenderList = append(appenderList, appender)
	lock.Unlock()
}

func notifyAppendLog(level int, msg string) {
	lock.Lock()
	for _, l := range appenderList {
		l.Log(level, msg)
	}
	lock.Unlock()
}
