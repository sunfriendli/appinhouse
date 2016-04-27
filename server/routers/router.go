package routers

import (
	. "appinhouse/server/constants"

	"appinhouse/server/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/:app/mobile/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/:app/mobile/list/:environment", &controllers.MainController{}, "get:List4Mobile")
	beego.Router("/api/:app/last", &controllers.MainController{}, "get:Last")
	beego.Router("/api/:app/list/:platform/:environment", &controllers.MainController{}, "get:List")
	//plist的url需要入库，所以固定下来，改的时候需要注意
	beego.Router(Url_Start+"/:app/"+Plist_Url_Uniform+"/:environment", &controllers.MainController{}, "get:PList")
	beego.Router("/api/:app/delete/:platform/:environment", &controllers.MainController{}, "delete:Delete")
	beego.Router("/api/:app/desc/:platform/:environment", &controllers.MainController{}, "post:Desc")
}
