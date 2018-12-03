package bases

import (
	"myGoApi/helpers"
	"strconv"
	"time"
)

type DebugResult struct {
	BaseResult
	StartTime string
	EndTime   string
	LogId     string
}

func (this *DebugResult) Init() {
	this.EndTime = strconv.FormatInt(time.Now().Unix(), 10)
	this.LogId = helpers.GetLoggerHelperObj()
}

func (this *DebugResult) SetData(data interface{}) {
	this.Data = data
}
