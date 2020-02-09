package log_extend

import (
	"sync"
)

var (
	appenderList = make([]Appender, 0)
	lock         = &sync.Mutex{}
)

const (
	Trace = 8
	Debug = 7
	Info  = 6
	Warn  = 4
	Error = 3
)

type Appender interface {
	Log(level int, msg string)
}

func AddLogAppender(appender Appender) {
	lock.Lock()
	appenderList = append(appenderList, appender)
	lock.Unlock()
}

func Log(level int, msg string) {
	lock.Lock()
	for _, l := range appenderList {
		l.Log(level, msg)
	}
	lock.Unlock()
}
