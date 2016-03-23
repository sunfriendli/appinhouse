package main

import (
	. "appinhouse/constants"
	_ "appinhouse/routers"

	"appinhouse/controllers"

	"github.com/astaxie/beego"
)

func main() {
	Root_Dir = beego.AppConfig.String("users::rootdir")
	Ios_Channel = beego.AppConfig.String("users::ioschannel")
	Domain = beego.AppConfig.String("users::domain")
	controllers.InitDirectory()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"appinhouse.log"}`)
	beego.SetStaticPath(Android_Release_Static_Path, controllers.GetDownDirectory(Android, Release))
	beego.SetStaticPath(Android_Dev_Static_Path, controllers.GetDownDirectory(Android, Dev))
	beego.SetStaticPath(Ios_Release_Static_Path, controllers.GetDownDirectory(Ios, Release))
	beego.SetStaticPath(Ios_Dev_Static_Path, controllers.GetDownDirectory(Ios, Dev))
	beego.Run()
}
