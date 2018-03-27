package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"

	"github.com/astaxie/beego"
)

type DeleteAppController struct {
	BaseController
}

func (c *DeleteAppController) DeleteApp() {
	dto := NewSuccessResponseDto()
	app := c.Ctx.Input.Param(":app")
	if app == "" || len(app) > App_Name_Len {
		beego.Info("DelApp param name  error !name:", app)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	err := models.AppListDao.Remove(app)
	if err != nil {
		beego.Info("DelApp  remove applist error !name:", app, "error:", err.Error())
		c.setError4Dto(err, dto)
		return
	}
	err = models.AppDao.Remove(app)
	if err != nil {
		beego.Info("DelApp  remove app error !name:", app, "error:", err.Error())
		c.setError4Dto(err, dto)
		return
	}

	c.Data["json"] = dto
	c.ServeJSON()
}
