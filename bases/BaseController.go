package bases

import (
	"github.com/astaxie/beego"
	"myGoApi/helpers"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["StartTime"] = time.Now()
}

func (o *BaseController) Success(data interface{}) {
	debugHelper := new(helpers.DebugHelper)
	var res IResult
	if debugHelper.IsDebug() {
		res = &DebugResult{}
	} else {
		res = &BaseResult{ErrNo: "0", ErrMsg: ""}
	}
	res.Init()
	res.SetData(data)
	o.Data["json"] = res
	o.ServeJSON()
}
