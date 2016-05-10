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

func (c *MainController) Apps() {

	dto := NewSuccessAppsResponseDto()
	page, err := c.GetInt("page", 1)
	if err != nil || page <= 0 || page > Max_Page {
		beego.Info("GetApps param page  error !page:", page)
		c.setError4Dto(ErrorParam, dto)
		return
	}
	count, err := models.AppDao.Count()
	if err != nil {
		beego.Error("GetApps count  error :" + err.Error())
		c.setError4Dto(ErrorParam, dto)
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
	ret, err := models.AppDao.GetList(start, end)
	if err != nil {
		c.setError4Dto(ErrorPageOut, dto)
		return
	}
	dto.Items = ret
	dto.Page = page
	dto.TotalPage = util.GetTotalPage(count, Page_Size)

	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) AddApp() {
	dto := NewSuccessResponseDto()
	app := c.Ctx.Input.Param(":app")
	if app == "" || len(app) > App_Name_Len {
		beego.Info("AddApp param name  error !name:", app)
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
		err = models.AppDao.Save(app)
		if err != nil {
			beego.Info("AddApp save app  error !name:", app, "error:", err.Error())
			c.setError4Dto(ErrorParam, dto)
			return
		}
	} else {
		c.setError4Dto(ErrorAppExistError, dto)
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

	err := models.AppDao.Remove(app)
	if err != nil {
		beego.Info("DelApp  remove app error !name:", app, "error:", err.Error())
		c.setError4Dto(ErrorParam, dto)
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
	osType := util.CheckAgent(userAgent)
	ret := make([]*ItemDto, 0, 4)
	switch osType {
	case util.OS_ANDROID:
		androiddev, err := getLastOne(Android, Dev, app)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androidre, err := getLastOne(Android, Release, app)
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
		iosdev, err := getLastOne(Ios, Dev, app)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release, app)
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
		iosdev, err := getLastOne(Ios, Dev, app)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release, app)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androiddev, err := getLastOne(Android, Dev, app)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androidre, err := getLastOne(Android, Release, app)
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
	page, err := c.GetInt("page", 1)
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
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)

	ret, total, err := getList(platform, env, start, end, app)

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
	page, err := c.GetInt("page", 1)
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
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)
	userAgent := c.Ctx.Request.UserAgent()
	osType := util.CheckAgent(userAgent)
	var ret []*ItemDto
	total := 0
	switch osType {
	case util.OS_ANDROID:
		ret, total, err = getList(Android, env, start, end, app)
		break
	case util.OS_IOS:
		ret, total, err = getList(Ios, env, start, end, app)
		break
	default:
		ret, total, err = getList(Android, env, start, end, app)
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
	score, err := time.Parse(time.RFC3339, ctime)
	if len(ctime) != len(F_Datetime) || err != nil {
		beego.Info("Desc error ,time:", ctime)
		c.setError4Dto(ErrorTimeFormat, dto)
		return
	}
	if platform == Ios && channel == Ios_Channel {
		plistUrl, err := c.savePlist(env, app, channel, version, id, title, softwareUrl, fullUrl, displayUrl)
		if err != nil {
			c.setError4Dto(err, dto)
			return
		}
		softwareUrl = plistUrl
	}
	info := new(models.DescInfo)
	info.Channel = channel
	info.Description = description
	info.Version = version
	info.Time = ctime
	info.Url = url
	info.SoftwareUrl = softwareUrl
	err = models.DescDao.Save(platform, env, app, info)
	if err != nil {
		c.setError4Dto(err, dto)
		return
	}
	err = models.DescListDao.Save(platform, env, app, info.Version, score)
	if err != nil {
		c.setError4Dto(err, dto)
		return
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) savePlist(env Environment, app, channel, version, id, title, softwareUrl, fullUrl, displayUrl string) (string, error) {

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

	err = models.PlistDao.Save(env, app, version, plist)
	if err != nil {
		beego.Info("PList CreatPlist fail :", err.Error())
		return "", err
	}
	return c.getPlistUrl(env, app, version), nil
}
func (c *MainController) getPlistUrl(env Environment, app, version string) string {
	envstr := Dev_Str
	if env == Release {
		envstr = Release_Str
	}
	return c.URLFor("MainController.PList", ":app", app, ":environment", envstr, ":version", version+Plist)
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
			models.DescDao.Remove(platform, env, app, versions)
			if platform == Ios {
				err := models.PlistDao.Remove(env, app, versions)
				if err != nil {
					c.setError4Dto(err, dto)
					return
				}
			}
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
func getLastOne(platform Platform, environment Environment, app string) (*ItemDto, error) {
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
	return converInfoTOItem(info, platform, environment, app), nil
}
func converInfoTOItem(info *models.DescInfo, platform Platform, environment Environment, app string) *ItemDto {

	item := new(ItemDto)
	switch platform {
	case Android:
		item.Platform = Android_Str
		item.Down = info.SoftwareUrl
		break
	case Ios:
		item.Platform = Ios_Str
		if info.Channel == Ios_Channel {
			item.Down = Https + Domain + info.SoftwareUrl
		} else {
			item.Down = info.SoftwareUrl
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

	t, _ := time.Parse(time.RFC3339, info.Time)

	rs := []rune(t.Local().String())
	item.Time = string(rs[0 : len(F_Datetime)-1])

	item.Url = info.Url
	item.Version = info.Version

	return item
}

func getList(platform Platform, environment Environment, start, end int, app string) ([]*ItemDto, int, error) {

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
		items = append(items, converInfoTOItem(info, platform, environment, app))
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
