package routers

import (
	"appinhouse/server/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/:app/mobile/last", &controllers.GetLastBuildsController{}, "get:GetLastBuilds")
	beego.Router("/api/:app/mobile/list/:environment", &controllers.GetBuildsController{}, "get:GetBuilds4Mobile")
	beego.Router("/api/:app/last", &controllers.GetLastBuildsController{}, "get:GetLastBuilds")
	beego.Router("/api/:app/list/:platform/:environment", &controllers.GetBuildsController{}, "get:GetBuilds")
	beego.Router("/api/:app/plist/:environment/:version", &controllers.GetPlistController{}, "get:GetPList")
	beego.Router("/api/:app/delete/:platform/:environment", &controllers.RemoveBuildsontroller{}, "delete:Clean")
	beego.Router("/api/:app/desc/:platform/:environment", &controllers.AddBuildController{}, "post:AddBuild")
	beego.Router("/api/:app/create", &controllers.CreateAppController{}, "post:CreateApp")
	beego.Router("/api/:app/update", &controllers.ModifyAppController{}, "post:ModifyApp")
	beego.Router("/api/:app/delete", &controllers.DeleteAppController{}, "delete:DeleteApp")
	beego.Router("/api/:app/get", &controllers.GetAppController{}, "get:GetApp")
	beego.Router("/api/apps", &controllers.GetAppsController{}, "get:GetApps")
	beego.Router("/api/modify/app", &controllers.ModifyDataController{}, "get:ModifyAppData")
	beego.Router("/api/move/:app/:operation", &controllers.MoveAppController{}, "post:MoveApp")

	beego.Router("/api/:app/:platform/:environment/:version", &controllers.GetBuildController{}, "get:GetBuild")

}
