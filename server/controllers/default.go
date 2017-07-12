package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"
	"appinhouse/server/util"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

func InitLogDirectory() {
	createFile(Log_Dir)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	dto := NewSuccessItemResponseDto()
	app := c.Ctx.Input.Param(":app")
	version := c.GetString(":version")
	if app == "" || len(app) > App_Name_Len {
		beego.Info("Get param name  error !name:", app)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	has, err := models.AppDao.Exist(app)
	if err != nil {
		beego.Info("Get Exist app  error !name:", app, "error:", err.Error())
		c.setError4Dto(ErrorParam, dto)
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("Get param platform  error !platform:", platformParam)
		c.setError4Dto(err, dto)
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("Get param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	if version == "" {
		beego.Info("Get get version  error !")
		c.setError4Dto(ErrorParam, dto)
		return
	}
	toffset := 0
	toffsetStr := c.GetString("time_offset")
	if toffsetStr == "" {
		toffset = Not_Offset
	} else {
		toffset, err = c.GetInt("time_offset")
		if err != nil {
			beego.Info("Get param toffset  error !time_offset:", toffset, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
	}
	if has {
		info, err := models.DescDao.Get(platform, env, app, version)
		if err != nil {
			beego.Info("Get get app  error !name:", app, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}

		dto.Item = converInfoTOItem(info, platform, env, app, toffset)

	} else {
		c.setError4Dto(ErrorAppNotExistError, dto)
		return
	}

	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) App() {
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
		dto.Item = converAppTOItem(app)
	} else {
		c.setError4Dto(ErrorAppNotExistError, dto)
		return
	}

	c.Data["json"] = dto
	c.ServeJSON()
}

func (c *MainController) Apps() {

	dto := NewSuccessAppsResponseDto()
	page, err := c.GetInt("page")
	if err != nil || page <= 0 || page > Max_Page {
		beego.Info("GetApps param page  error !page:", page)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	count, err := models.AppListDao.Count()
	if err != nil {
		beego.Error("GetApps count  error :" + err.Error())
		c.setError4Dto(err, dto)
		return
	}

	if count == 0 {
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)
	if start > count {
		c.setError4Dto(ErrorPageOut, dto)
		return
	}
	if end > count {
		end = count - 1
	}
	apps, err := models.AppListDao.GetList(start, end)
	if err != nil {
		c.setError4Dto(err, dto)
		return
	}
	if apps != nil {
		ret, err := models.AppDao.MGet(apps)
		if err != nil {
			c.setError4Dto(err, dto)
			return
		}
		items := make([]*AppDto, 0, len(ret))
		for _, info := range ret {
			items = append(items, converAppTOItem(info))
		}
		dto.Items = items
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(count, Page_Size)
	}

	c.Data["json"] = dto
	c.ServeJSON()
}

func (c *MainController) AddApp() {
	dto := NewSuccessResponseDto()
	app := c.Ctx.Input.Param(":app")
	desc := c.GetString("description")
	alias := c.GetString("alias")

	if app == "" || len(app) > App_Name_Len || desc == "" {
		beego.Info("AddApp param name  error !name:", app, "desc:", desc)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	has, err := models.AppDao.Exist(app)
	if err != nil {
		beego.Info("AddApp Exist app  error !name:", app, "error:", err.Error())
		c.setError4Dto(ErrorParam, dto)
		return
	}
	if !has {
		appinfo := new(models.AppInfo)
		appinfo.App = app
		appinfo.Description = desc
		appinfo.Alias = alias
		err = models.AppDao.Save(appinfo)
		if err != nil {
			beego.Info("AddApp save app  error !name:", app, "error:", err.Error())
			c.setError4Dto(err, dto)
			return
		}

		err = models.AppListDao.Save(app)
		if err != nil {
			beego.Info("AddApp save applist  error !name:", app, "error:", err.Error())
			c.setError4Dto(err, dto)
			return
		}

	} else {
		c.setError4Dto(ErrorAppExistError, dto)
		return
	}

	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) ModifyApp() {
	dto := NewSuccessResponseDto()
	app := c.Ctx.Input.Param(":app")
	desc := c.GetString("description")
	alias := c.GetString("alias")

	if app == "" || len(app) > App_Name_Len || desc == "" {
		beego.Info("ModifyApp param name  error !name:", app, "desc:", desc)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	has, err := models.AppDao.Exist(app)
	if err != nil {
		beego.Info("ModifyApp Exist app  error !name:", app, "error:", err.Error())
		c.setError4Dto(ErrorParam, dto)
		return
	}
	if has {
		appinfo := new(models.AppInfo)
		appinfo.App = app
		appinfo.Description = desc
		appinfo.Alias = alias
		err = models.AppDao.Save(appinfo)
		if err != nil {
			beego.Info("AddApp save app  error !name:", app, "error:", err.Error())
			c.setError4Dto(err, dto)
			return
		}
	} else {
		c.setError4Dto(ErrorAppNotExistError, dto)
		return
	}

	c.Data["json"] = dto
	c.ServeJSON()
}

func (c *MainController) DelApp() {
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
func (c *MainController) Last() {
	dto := NewSuccessItemsResponseDto()
	userAgent := c.Ctx.Request.UserAgent()
	appParam := c.Ctx.Input.Param(":app")

	app, err := appExist(appParam)
	if err != nil {
		beego.Info("Last param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}
	toffset := 0
	toffsetStr := c.GetString("time_offset")
	if toffsetStr == "" {
		toffset = Not_Offset
	} else {
		toffset, err = c.GetInt("time_offset")
		if err != nil {
			beego.Info("Last param toffset  error !time_offset:", toffset, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
	}
	osType := util.CheckAgent(userAgent)
	ret := make([]*ItemDto, 0, 4)
	switch osType {
	case util.OS_ANDROID:
		androiddev, err := getLastOne(Android, Dev, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androidre, err := getLastOne(Android, Release, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		if androiddev != nil {
			ret = append(ret, androiddev)
		}
		if androidre != nil {
			ret = append(ret, androidre)
		}
		break
	case util.OS_IOS:
		iosdev, err := getLastOne(Ios, Dev, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		if iosdev != nil {
			ret = append(ret, iosdev)
		}
		if iosre != nil {
			ret = append(ret, iosre)
		}
		break
	default:
		iosdev, err := getLastOne(Ios, Dev, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androiddev, err := getLastOne(Android, Dev, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androidre, err := getLastOne(Android, Release, app, toffset)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}

		if androiddev != nil {
			ret = append(ret, androiddev)
		}
		if androidre != nil {
			ret = append(ret, androidre)
		}

		if iosdev != nil {
			ret = append(ret, iosdev)
		}
		if iosre != nil {
			ret = append(ret, iosre)
		}
		break
	}
	dto.Msg = GetMsg(dto.Code)
	dto.Items = ret
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) List() {
	dto := NewSuccessItemsResponsePageDto()
	page, err := c.GetInt("page")
	if err != nil || page <= 0 || page > Max_Page {
		beego.Info("List param page  error !page:", page)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	appParam := c.Ctx.Input.Param(":app")
	app, err := appExist(appParam)
	if err != nil {
		beego.Info("List param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("List param platform  error !platform:", platformParam)
		c.setError4Dto(err, dto)
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("List param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	toffset := 0
	toffsetStr := c.GetString("time_offset")
	if toffsetStr == "" {
		toffset = Not_Offset
	} else {
		toffset, err = c.GetInt("time_offset")
		if err != nil {
			beego.Info("List param toffset  error !time_offset:", toffset, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
	}
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)

	ret, total, err := getList(platform, env, start, end, app, toffset)

	if err != nil {
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
	} else {
		dto.Items = ret
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(total, Page_Size)
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) List4Mobile() {
	dto := NewSuccessItemsResponsePageDto()
	page, err := c.GetInt("page")
	if err != nil || page <= 0 || page > Max_Page {
		beego.Info("List4Mobile param page  error !page:", page)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	appParam := c.Ctx.Input.Param(":app")
	app, err := appExist(appParam)
	if err != nil {
		beego.Info("List4Mobile param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("List4Mobile param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	toffset := 0
	toffsetStr := c.GetString("time_offset")
	if toffsetStr == "" {
		toffset = Not_Offset
	} else {
		toffset, err = c.GetInt("time_offset")
		if err != nil {
			beego.Info("List4Mobile param toffset  error !time_offset:", toffset, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
	}
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)
	userAgent := c.Ctx.Request.UserAgent()
	osType := util.CheckAgent(userAgent)
	var ret []*ItemDto
	total := 0
	switch osType {
	case util.OS_ANDROID:
		ret, total, err = getList(Android, env, start, end, app, toffset)
		break
	case util.OS_IOS:
		ret, total, err = getList(Ios, env, start, end, app, toffset)
		break
	default:
		ret, total, err = getList(Android, env, start, end, app, toffset)
		break
	}

	if err != nil {
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
	} else {
		dto.Items = ret
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(total, Page_Size)
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) PList() {
	dto := NewSuccessResponseDto()
	version := c.Ctx.Input.Param(":version")
	version = strings.Replace(version, Plist, "", -1)
	if version == "" {
		beego.Info("GetPList Params fail .version:", version)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := appExist(appParam)
	if err != nil {
		beego.Info("GetPList param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("GetPList param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	plist, err := models.PlistDao.Get(env, app, version)
	if err != nil {
		beego.Info("GetPList from redis fail :", err.Error())
		c.setError4Dto(err, dto)
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/x-plist")
	c.Ctx.Output.Body([]byte(plist))
}
func (c *MainController) Desc() {
	dto := NewSuccessResponseDto()

	version := c.GetString("version")
	ctime := c.GetString("time")
	description := c.GetString("description")
	url := c.GetString("url")
	channel := c.GetString("channel")
	softwareUrl := c.GetString("software_url")
	//当生成ios描述文件时，生成plist需要
	id := c.GetString("id")
	title := c.GetString("title")
	fullUrl := c.GetString("full_url")
	displayUrl := c.GetString("display_url")
	extendSoftwareUrlName := c.GetString("software_url_extend_name")

	if ctime == "" || version == "" || channel == "" || softwareUrl == "" {
		beego.Info("Desc Params fail version:", version, "time:", ctime, "description:", description,
			"software_url", softwareUrl, "channel:", channel)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := appExist(appParam)
	if err != nil {
		beego.Info("Desc param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("Desc param platform  error !platform:", platformParam)
		c.setError4Dto(err, dto)
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("Desc param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}

	if !existHttpPath(softwareUrl) {
		beego.Info("Desc error archive file not exsit,url:", softwareUrl)
		c.setError4Dto(ErrorIPANotExistError, dto)
		return
	}
	score, err := time.Parse(ISO_Datetime, ctime)
	if err != nil {
		beego.Info("Desc error ,time:", ctime)
		c.setError4Dto(ErrorTimeFormat, dto)
		return
	}
	if platform == Ios && channel == Ios_Channel {
		plistUrl, err := c.savePlist(env, app, channel, version, extendSoftwareUrlName, id, title, softwareUrl, fullUrl, displayUrl)
		if err != nil {
			c.setError4Dto(err, dto)
			return
		}
		softwareUrl = plistUrl
	}
	info, err := models.DescDao.Get(platform, env, app, version)
	if err != nil {
		beego.Info("get Desc error !platform:", platform, "env:", env, "app:", app, "version:", version, "error:", err.Error())
		c.setError4Dto(ErrorDB, dto)
		return
	}
	newInfo := false
	if info == nil {
		newInfo = true
		info = new(models.DescInfo)
		info.Channel = channel
		info.Description = description
		info.Version = version
		info.Time = ctime
		info.Url = url
	}

	if extendSoftwareUrlName != "" {
		if info.ExtendSoftwareUrls == nil {
			info.ExtendSoftwareUrls = make(map[string]string)
		}
		info.ExtendSoftwareUrls[extendSoftwareUrlName] = softwareUrl
	} else {
		info.SoftwareUrl = softwareUrl
	}

	err = models.DescDao.Save(platform, env, app, info)
	if err != nil {
		c.setError4Dto(err, dto)
		return
	}
	if newInfo {
		err = models.DescListDao.Save(platform, env, app, info.Version, score)
		if err != nil {
			c.setError4Dto(err, dto)
			return
		}
	}

	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) savePlist(env Environment, app, channel, version, extendSoftwareUrlName, id, title, softwareUrl, fullUrl, displayUrl string) (string, error) {

	if id == "" || title == "" {
		beego.Info("PList Params fail .id:", id, "title:", title)
		return "", ErrorParam
	}
	hasFile := true
	if fullUrl != "" {
		if !existHttpPath(fullUrl) {
			hasFile = false
		}
	}
	if displayUrl != "" {
		if !existHttpPath(displayUrl) {
			hasFile = false
		}
	}
	if !hasFile {
		beego.Info("PList error fullUrl or displayUrl not exsit.fullUrl:", fullUrl, "displayUrl", displayUrl)
		return "", ErrorParam
	}
	plist, err := models.NewPlist(id, version, title, softwareUrl, fullUrl, displayUrl)

	if err != nil {
		beego.Info("PList NewPlist error :", err.Error())
		return "", err
	}

	err = models.PlistDao.Save(env, app, version, extendSoftwareUrlName, plist)
	if err != nil {
		beego.Info("PList CreatPlist fail :", err.Error())
		return "", err
	}
	return c.getPlistUrl(env, app, version, extendSoftwareUrlName), nil
}
func (c *MainController) getPlistUrl(env Environment, app, version, extendSoftwareUrlName string) string {
	envstr := Dev_Str
	if env == Release {
		envstr = Release_Str
	}
	plist := models.PlistDao.GetField(version, extendSoftwareUrlName) + Plist
	return c.URLFor("MainController.PList", ":app", app, ":environment", envstr, ":version", plist)
}
func (c *MainController) Delete() {
	dto := NewSuccessResponseDto()
	residue, err := c.GetInt("residue", -1)
	if err != nil || residue < Min_Residue {
		beego.Info("Delete param residue  error !residue:", residue)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	appParam := c.Ctx.Input.Param(":app")
	app, err := appExist(appParam)
	if err != nil {
		beego.Info("Delete param app  error !app:", appParam)
		c.setError4Dto(err, dto)
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("Delete param platform  error !platform:", platformParam)
		c.setError4Dto(err, dto)
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("Delete param environment  error !environment:", envParam)
		c.setError4Dto(err, dto)
		return
	}
	count, err := models.DescListDao.Count(platform, env, app)
	if err != nil {
		beego.Error("Delete count  error :" + err.Error())
		c.setError4Dto(err, dto)
		return
	}

	if residue < count {
		versions, err := models.DescListDao.Remove(platform, env, app, residue)
		if err != nil {
			c.setError4Dto(err, dto)
			return
		}

		if versions != nil && len(versions) > 0 {

			if platform == Ios {
				extendUrls := make([]string, 10)
				infos, _ := models.DescDao.MGet(platform, env, app, versions)
				for _, info := range infos {
					if info.ExtendSoftwareUrls != nil && len(info.ExtendSoftwareUrls) > 0 {

						for k, _ := range info.ExtendSoftwareUrls {
							extendUrls = append(extendUrls, models.PlistDao.GetField(info.Version, k))
						}
					}
				}
				err := models.PlistDao.Remove(env, app, extendUrls)
				err = models.PlistDao.Remove(env, app, versions)
				if err != nil {
					c.setError4Dto(err, dto)
					return
				}
			}
			models.DescDao.Remove(platform, env, app, versions)
		}
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) setError4Dto(err error, dto ResponseDto) {
	code := GetErrCode(err)
	dto.SetCode(code)
	dto.SetMsg(GetMsg(code))
	c.Data["json"] = dto
	c.ServeJSON()
}
func getLastOne(platform Platform, environment Environment, app string, toffset int) (*ItemDto, error) {
	version, err := models.DescListDao.GetLast(platform, environment, app)
	if err != nil {
		beego.Error("getLastOne  GetLast error:" + err.Error())
		return nil, err
	}
	if version == "" {
		return nil, nil
	}
	info, err := models.DescDao.Get(platform, environment, app, version)
	if err != nil {
		beego.Error("getLastOne  Get error:" + err.Error())
		return nil, err
	}
	return converInfoTOItem(info, platform, environment, app, toffset), nil
}
func converInfoTOItem(info *models.DescInfo, platform Platform, environment Environment, app string, toffset int) *ItemDto {
	if info == nil {
		return nil
	}
	item := new(ItemDto)
	item.ExtendUrls = make(map[string]string)

	switch platform {
	case Android:
		item.Platform = Android_Str
		item.Down = info.SoftwareUrl
		if info.ExtendSoftwareUrls != nil && len(info.ExtendSoftwareUrls) > 0 {
			for k, v := range info.ExtendSoftwareUrls {
				item.ExtendUrls[k] = v
			}
		}
		break
	case Ios:
		item.Platform = Ios_Str
		if info.Channel == Ios_Channel {
			if info.SoftwareUrl != "" {
				item.Down = Https + Domain + info.SoftwareUrl
			}
			if info.ExtendSoftwareUrls != nil && len(info.ExtendSoftwareUrls) > 0 {
				for k, v := range info.ExtendSoftwareUrls {
					item.ExtendUrls[k] = Https + Domain + v
				}
			}
		} else {
			item.Down = info.SoftwareUrl
			if info.ExtendSoftwareUrls != nil && len(info.ExtendSoftwareUrls) > 0 {
				for k, v := range info.ExtendSoftwareUrls {
					item.ExtendUrls[k] = v
				}
			}
		}
		break
	}
	switch environment {
	case Release:
		item.Env = Release_Str
		break
	case Dev:
		item.Env = Dev_Str
		break
	}
	item.Channel = info.Channel
	item.Description = info.Description

	item.Time = converToClientTime(info.Time, toffset)

	item.Url = info.Url
	item.Version = info.Version

	return item
}
func converToClientTime(ctime string, toffset int) string {
	l := len(ctime)
	if toffset == Not_Offset {
		if len(ISO_Datetime) == l {
			t, _ := time.Parse(ISO_Datetime, ctime)
			rs := []rune(t.Local().String())
			return string(rs[0 : len(ISO_Datetime)-5])
		} else if len(UTC_Datetime) == l {
			t, _ := time.Parse(time.RFC3339, ctime)
			rs := []rune(t.Local().String())
			return string(rs[0 : len(UTC_Datetime)-1])
		} else {
			return ctime
		}

	} else {
		if len(ISO_Datetime) == l {
			return util.FixUTCTimeToOffsetSpecifiedZoneTime(toffset*60, ctime, ISO_Datetime, View_Datetime)
		} else if len(UTC_Datetime) == l {
			return util.FixUTCTimeToOffsetSpecifiedZoneTime(toffset*60, ctime, UTC_Datetime, View_Datetime)
		} else {
			return ctime
		}
	}

}
func converAppTOItem(info *models.AppInfo) *AppDto {

	item := new(AppDto)
	item.App = info.App
	item.Desc = info.Description
	item.Alias = info.Alias
	return item
}
func getList(platform Platform, environment Environment, start, end int, app string, toffset int) ([]*ItemDto, int, error) {

	count, err := models.DescListDao.Count(platform, environment, app)
	if err != nil {
		beego.Error("getList count  error :" + err.Error())
		return nil, 0, err
	}
	if count == 0 {
		return make([]*ItemDto, 0, 0), 0, nil
	}
	if start > count {
		beego.Info("getList call page start greate than size")
		return nil, 0, ErrorPageOut
	}
	if end > count {
		end = count - 1
	}

	versions, err := models.DescListDao.GetList(platform, environment, app, start, end)

	if err != nil {
		beego.Error("getList from redis error :" + err.Error())
		return nil, 0, err
	}
	size := len(versions)
	if versions == nil || size == 0 {
		return make([]*ItemDto, 0, size), 0, nil
	}
	items := make([]*ItemDto, 0, size)
	infos, err := models.DescDao.MGet(platform, environment, app, versions)
	for _, info := range infos {
		items = append(items, converInfoTOItem(info, platform, environment, app, toffset))
	}

	return items, count, nil
}
func appExist(app string) (string, error) {
	if app == "" {
		return "", ErrorParam
	}
	has, err := models.AppDao.Exist(app)
	if err != nil {
		return app, err
	}
	if !has {
		return app, ErrorAppNotExistError
	}
	return app, nil
}
func createFile(dir string) (string, error) {
	src := dir
	if isExist(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0757); err != nil {
		if os.IsPermission(err) {
			beego.Info("permission denied path:" + src)
			panic("permission denied")
		}
		return "", err
	}

	return src, nil
}
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//判断服务器上的文件是否存在
func existHttpPath(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		beego.Info("existHttpPath error url:", url, " error:", err.Error())
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}
