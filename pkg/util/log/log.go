// Copyright 2016 fatedier, fatedier@gmail.com
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

package log

import (
	"fmt"
	"github.com/fatedier/beego/logs"
	"github.com/fatedier/frp/sharefrp/pkg/logx"
)

// Log is the under logx object
var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(200)
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(Log.GetLogFuncCallDepth() + 1)
}

func InitLog(logWay string, logFile string, logLevel string, maxdays int64, disableLogColor bool) {
	SetLogFile(logWay, logFile, maxdays, disableLogColor)
	SetLogLevel(logLevel)
}

// SetLogFile to configure logx params
// logWay: file or console
func SetLogFile(logWay string, logFile string, maxdays int64, disableLogColor bool) {
	if logWay == "console" {
		params := ""
		if disableLogColor {
			params = fmt.Sprintf(`{"color": false}`)
		}
		Log.SetLogger("console", params)
	} else {
		params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, logFile, maxdays)
		Log.SetLogger("file", params)
	}
}

// SetLogLevel set logx level, default is warning
// value: error, warning, info, debug, trace
func SetLogLevel(logLevel string) {
	level := 4 // warning
	switch logLevel {
	case "error":
		level = 3
	case "warn":
		level = 4
	case "info":
		level = 6
	case "debug":
		level = 7
	case "trace":
		level = 8
	default:
		level = 4
	}
	Log.SetLevel(level)
}

// wrap logx

func Error(format string, v ...interface{}) {
	Log.Error(format, v...)
	logx.Log(logx.Error, fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	Log.Warn(format, v...)
	logx.Log(logx.Warn, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	Log.Info(format, v...)
	logx.Log(logx.Info, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	Log.Debug(format, v...)
	logx.Log(logx.Debug, fmt.Sprintf(format, v...))
}

func Trace(format string, v ...interface{}) {
	Log.Trace(format, v...)
	logx.Log(logx.Trace, fmt.Sprintf(format, v...))
}
