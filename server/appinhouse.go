package main

import (
	. "appinhouse/server/constants"

	_ "appinhouse/server/routers"

	"appinhouse/server/controllers"

	"strings"

	"appinhouse/server/models"

	"github.com/astaxie/beego"

	"github.com/Unknwon/goconfig"
)

func main() {
	setLog()
	setParam()
	models.InitValue()
	beego.Run()
}
func setLog() {
	Log_Dir = beego.AppConfig.String("users::log_dir")
	controllers.InitLogDirectory()
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"`+Log_Dir+Log_File+`"}`)

}
func setParam() {

	Ios_Channel = beego.AppConfig.String("users::ios_channel")
	beego.Info("app.conf-> Ios_Channel:", Ios_Channel)
	if Ios_Channel == "" {
		panic("app.conf not have users::ios_channel ")
	}
	Domain = beego.AppConfig.String("users::domain")
	beego.Info("app.conf-> Domain:", Domain)
	if Domain == "" {
		panic("app.conf not have users::domain ")
	}
	apps := beego.AppConfig.String("users::app_names")
	beego.Info("app.conf-> app_names:", apps)
	if apps == "" {
		panic("app.conf not have users::app_names ")
	}
	AddApps(strings.Split(apps, ";"))
	Full_Size_Image = beego.AppConfig.String("users::full_size_image")
	beego.Info("app.conf-> Full_Size_Image:", Full_Size_Image)
	if Full_Size_Image == "" {
		panic("app.conf not have users::full_size_image ")
	}
	Display_Image = beego.AppConfig.String("users::display_image")
	beego.Info("app.conf-> Display_Image:", Display_Image)
	if Display_Image == "" {
		panic("app.conf not have users::display_image ")
	}
	Min_Residue, _ = beego.AppConfig.Int("users::min_residue")
	beego.Info("app.conf-> Min_Residue:", Min_Residue)
	if Min_Residue == 0 {
		panic("app.conf not have users::min_residue or not int")
	}
	Page_Size, _ = beego.AppConfig.Int("users::page_size")
	beego.Info("app.conf-> Page_Size:", Page_Size)
	if Page_Size == 0 {
		panic("app.conf not have users::page_size or not int")
	}
	Max_Page, _ = beego.AppConfig.Int("users::max_page")
	beego.Info("app.conf-> Max_Page:", Max_Page)
	if Max_Page == 0 {
		panic("app.conf not have users::max_page or not int")
	}

	confDir := beego.AppConfig.String("redis::conf_dir")
	beego.Info("app.conf-> conf_dir:", confDir)
	if confDir == "" {
		panic("app.conf not have users::conf_dir ")
	}

	addr := beego.AppConfig.String("redis::env_addr_name")
	beego.Info("app.conf-> env_addr_name:", addr)
	if addr == "" {
		panic("app.conf not have users::env_addr_name ")
	}

	pwd := beego.AppConfig.String("redis::env_password_name")
	beego.Info("app.conf-> env_password_name:", pwd)
	if pwd == "" {
		panic("app.conf not have users::env_password_name ")
	}

	getRedisConf(confDir, addr, pwd)

	Redis_DB = beego.AppConfig.DefaultInt("redis::db", Redis_DB)
	beego.Info("app.conf-> db:", Redis_DB)

	Redis_PoolSize = beego.AppConfig.DefaultInt("redis::pool_siz", Redis_PoolSize)
	beego.Info("app.conf-> pool_siz:", Redis_PoolSize)
}
func getRedisConf(confDir, addr, pwd string) {
	c, err := goconfig.LoadConfigFile(confDir)
	if err != nil {
		panic("redis conf path error :" + err.Error())
	}
	Redis_Addr, err = c.GetValue("", addr)
	beego.Info("Redis_Addr:", Redis_Addr)
	if err != nil {
		panic("redis conf addr error :" + err.Error())
	}
	Redis_Password, err = c.GetValue("", pwd)
	beego.Info("Redis_Password:", Redis_Password)

}
