package bases

import (
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
	this.LogId = string("xxxxxxxxxxxxxxxxxxxxxxxxx")
}

func (this *DebugResult) SetData(data interface{}) {
	this.Data = data
}
