package main

import (
	"C"
	"container/list"
	sharefrp "github.com/fatedier/frp/sharefrp/mobile"
	"net/url"
	"strconv"
	"time"
)

var (
	listener          = &DefaultLogListener{}
	isInit            = false
	bufLogList        = list.New()
	bufLogListMaxSize = 1000
)

type DefaultLogListener struct {
}

func (listener *DefaultLogListener) Log(level int, msg string) {
	query := "time=" + strconv.FormatInt(time.Now().Unix(), 10) + "&level=" + strconv.Itoa(level) + "&msg=" + url.PathEscape(msg)
	if bufLogList.Len() >= bufLogListMaxSize {
		l := bufLogList.Front()
		if l != nil {
			_ = bufLogList.Remove(l)
		}
	}
	bufLogList.PushBack(query)
}

//export LOG_LEVEL_ERROR
func LOG_LEVEL_ERROR() C.int {
	return C.int(sharefrp.LOG_LEVEL_ERROR)
}

//export LOG_LEVEL_WARN
func LOG_LEVEL_WARN() C.int {
	return C.int(sharefrp.LOG_LEVEL_WARN)
}

//export LOG_LEVEL_INFO
func LOG_LEVEL_INFO() C.int {
	return C.int(sharefrp.LOG_LEVEL_INFO)
}

//export LOG_LEVEL_DEBUG
func LOG_LEVEL_DEBUG() C.int {
	return C.int(sharefrp.LOG_LEVEL_DEBUG)
}

//export LOG_LEVEL_TRACE
func LOG_LEVEL_TRACE() C.int {
	return C.int(sharefrp.LOG_LEVEL_TRACE)
}

//export STATUS_NONE
func STATUS_NONE() C.int {
	return C.int(sharefrp.STATUS_NONE)
}

//export STATUS_RUNNING
func STATUS_RUNNING() C.int {
	return C.int(sharefrp.STATUS_RUNNING)
}

//export STATUS_STOP
func STATUS_STOP() C.int {
	return C.int(sharefrp.STATUS_STOP)
}

//export Init
func Init(maxLogSize C.int) {
	bufLogListMaxSize = int(maxLogSize)
	if !isInit {
		sharefrp.ListenLog(listener)
		isInit = true
	}
}

//export NextLog
func NextLog() *C.char {
	l := bufLogList.Front()
	if l != nil {
		return C.CString(bufLogList.Remove(l).(string))
	}
	return C.CString("")
}

//export GetFrpVersion
func GetFrpVersion() *C.char {
	return C.CString(sharefrp.GetFrpVersion())
}

//export GetSharedFrpVersion
func GetSharedFrpVersion() *C.char {
	return C.CString(sharefrp.GetSharedFrpVersion())
}

func main() {}
