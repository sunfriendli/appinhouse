package main

import (
	. "appinhouse/server/constants"

	_ "appinhouse/server/routers"

	"appinhouse/server/controllers"

	"appinhouse/server/models"

	"appinhouse/server/filters"

	"github.com/astaxie/beego"
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
	setFilter()

}

func setFilter() {
	beego.InsertFilter("/api/:app/delete/:platform/:environment", beego.BeforeRouter, filters.SecretKeyFilter)
	beego.InsertFilter("/api/:app/desc/:platform/:environment", beego.BeforeRouter, filters.SecretKeyFilter)
	beego.InsertFilter("/api/:app/create", beego.BeforeRouter, filters.SecretKeyFilter)
	beego.InsertFilter("/api/:app/update", beego.BeforeRouter, filters.SecretKeyFilter)
	beego.InsertFilter("/api/:app/delete", beego.BeforeRouter, filters.SecretKeyFilter)
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
	App_Name_Len, _ = beego.AppConfig.Int("users::app_name_len")
	beego.Info("app.conf-> app_name_len:", App_Name_Len)
	if App_Name_Len == 0 {
		panic("app.conf not have users::app_name_len ")
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

	Redis_Addr = beego.AppConfig.String("redis::addr")
	beego.Info("app.conf-> addr:", Redis_Addr)
	if Redis_Addr == "" {
		panic("app.conf not have users::addr ")
	}

	Redis_Password = beego.AppConfig.String("redis::password")
	beego.Info("app.conf-> password:", Redis_Password)

	Redis_DB = beego.AppConfig.DefaultInt("redis::db", Redis_DB)
	beego.Info("app.conf-> db:", Redis_DB)

	Redis_PoolSize = beego.AppConfig.DefaultInt("redis::pool_siz", Redis_PoolSize)
	beego.Info("app.conf-> pool_siz:", Redis_PoolSize)

	Secret_Key = beego.AppConfig.String("authentication::secret_key")
	beego.Info("app.conf-> secret_key:", Secret_Key)
	if Secret_Key == "" {
		panic("app.conf not authentication::secret_key")
	}

	group := beego.AppConfig.Strings("securityGroup::inbound")
	beego.Info("app.conf-> inbound:", group)
	if group == nil || len(group) == 0 {
		panic("app.conf not securityGroup::inbound")
	}
	for _, value := range group {
		Security_Group[value] = true
	}
	beego.Info("app.conf-> inbound:", Security_Group)
}
