package helpers

import (
	"github.com/rs/xid"
)

var logId string

func init() {
	logId = xid.New().String()
}
func GetLoggerHelperObj() string {
	if "" == logId {
		logId = xid.New().String()
	}
	return logId
}

func ClearLogId() {
	logId = ""
}

type LogHelper struct {
	LogId string
}

func (l *LogHelper) SetLogId() {
	l.LogId = xid.New().String()
}

func (l *LogHelper) GetLogId() string {
	return l.LogId
}

func (l *LogHelper) FlushLogId() string {
	l.LogId = xid.New().String()
	return l.LogId
}