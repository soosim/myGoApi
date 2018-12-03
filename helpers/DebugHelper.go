package helpers

import "github.com/astaxie/beego"

func IsDebug() bool {
	return beego.AppConfig.String("runmode") == "dev"
}
