package routers

import (
	"appinhouse/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/api/mobile/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/mobile/list/dev", &controllers.MainController{}, "get:DevList")
	beego.Router("/api/mobile/list/release", &controllers.MainController{}, "get:ReleaseList")
	beego.Router("/api/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/list/android/dev", &controllers.MainController{}, "get:AndroidDevList")
	beego.Router("/api/list/android/release", &controllers.MainController{}, "get:AndroidReleaseList")
	beego.Router("/api/list/ios/dev", &controllers.MainController{}, "get:IosDevList")
	beego.Router("/api/list/ios/release", &controllers.MainController{}, "get:IosReleaseList")
}
