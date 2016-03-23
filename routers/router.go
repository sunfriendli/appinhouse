package routers

import (
	"appinhouse/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/v1/mobile/last", &controllers.MainController{}, "get:Last")
	beego.Router("/v1/mobile/list/dev", &controllers.MainController{}, "get:DevList")
	beego.Router("/v1/mobile/list/release", &controllers.MainController{}, "get:ReleaseList")
	beego.Router("/v1/last", &controllers.MainController{}, "get:Last")
	beego.Router("/v1/list/android/dev", &controllers.MainController{}, "get:AndroidDevList")
	beego.Router("/v1/list/android/release", &controllers.MainController{}, "get:AndroidReleaseList")
	beego.Router("/v1/list/ios/dev", &controllers.MainController{}, "get:IosDevList")
	beego.Router("/v1/list/ios/release", &controllers.MainController{}, "get:IosReleaseList")
}
