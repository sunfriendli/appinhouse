package routers

import (
	"appinhouse/server/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/:app/mobile/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/:app/mobile/list/:environment", &controllers.MainController{}, "get:List4Mobile")
	beego.Router("/api/:app/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/:app/list/:platform/:environment", &controllers.MainController{}, "get:List")
	beego.Router("/api/:app/plist/:environment/:version", &controllers.MainController{}, "get:PList")
	beego.Router("/api/:app/delete/:platform/:environment", &controllers.MainController{}, "delete:Delete")
	beego.Router("/api/:app/desc/:platform/:environment", &controllers.MainController{}, "post:Desc")
	beego.Router("/api/:app/create", &controllers.MainController{}, "post:AddApp")
	beego.Router("/api/:app/update", &controllers.MainController{}, "post:ModifyApp")
	beego.Router("/api/:app/delete", &controllers.MainController{}, "delete:DelApp")
	beego.Router("/api/:app/get", &controllers.MainController{}, "get:App")
	beego.Router("/api/apps", &controllers.MainController{}, "get:Apps")
	beego.Router("/api/modify/app", &controllers.ModifyDataController{}, "get:ModifyAppData")
	beego.Router("/api/move/:app/:operation", &controllers.MainController{}, "post:MoveApp")

	beego.Router("/api/:app/:platform/:environment/:version", &controllers.MainController{}, "get:Get")
}
