// getapp
package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"

	"github.com/astaxie/beego"
)

type GetAppController struct {
	BaseController
}

func (c *GetAppController) GetApp() {
	dto := NewSuccessAppResponseDto()
	app := c.Ctx.Input.Param(":app")

	if app == "" || len(app) > App_Name_Len {
		beego.Info("App param name  error !name:", app)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	has, err := models.AppDao.Exist(app)
	if err != nil {
		beego.Info("App Exist app  error !name:", app, "error:", err.Error())
		c.setError4Dto(ErrorParam, dto)
		return
	}
	if has {
		app, err := models.AppDao.Get(app)
		if err != nil {
			beego.Info("App get app  error !name:", app, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
		dto.Item = c.converAppTOItem(app)
	} else {
		c.setError4Dto(ErrorAppNotExistError, dto)
		return
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
