package helpers

import "github.com/astaxie/beego"

type DebugHelper struct {
	isDebug bool
}

func (this *DebugHelper) IsDebug() bool {
	return beego.AppConfig.String("runmode") == "dev"
}
