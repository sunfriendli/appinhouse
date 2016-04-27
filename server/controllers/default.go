package controllers

import (
	. "appinhouse/server/constants"
	"appinhouse/server/models"
	"appinhouse/server/util"
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/astaxie/beego"
)

func InitLogDirectory() {
	createFile(Log_Dir)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Last() {
	dto := NewSuccessItemsResponseDto()
	userAgent := c.Ctx.Request.UserAgent()
	appParam := c.Ctx.Input.Param(":app")

	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("Last param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
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
	page, _ := c.GetInt("page", 1)
	if page <= 0 || page > Max_Page {
		beego.Info("List param page  error !page:", page)
		dto.Code = ErrParams
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("List param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("List param platform  error !platform:", platformParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("List param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
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
	page, _ := c.GetInt("page", 1)
	if page <= 0 || page > Max_Page {
		beego.Info("List4Mobile param page  error !page:", page)
		dto.Code = ErrParams
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("List4Mobile param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("List4Mobile param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
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
	version := c.GetString("version")
	if version == "" {
		beego.Info("GetPList Params fail .version:", version)
		dto.Code = ErrParams
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("GetPList param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("GetPList param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	plist, err := models.PlistDao.Get(env, app, version)
	if err != nil {
		beego.Info("GetPList from redis fail :", err.Error())
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
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
	softwareUrl := c.GetString("softwareUrl")
	//当生成ios描述文件时，生成plist需要
	id := c.GetString("id")
	title := c.GetString("title")
	fullUrl := c.GetString("fullUrl")
	displayUrl := c.GetString("displayUrl")

	if ctime == "" || version == "" || url == "" || channel == "" || softwareUrl == "" {
		beego.Info("Desc Params fail version:", version, "time:", ctime, "description:", description,
			"url:", url, "softwareUrl", softwareUrl, "channel:", channel)
		dto.Code = ErrParams
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("Desc param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("Desc param platform  error !platform:", platformParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("Desc param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	if !existHttpPath(softwareUrl) {
		beego.Info("PList error archive file not exsit,url:", softwareUrl)
		dto.Code = ErrIPANotExistError
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	score, err := time.Parse(F_Datetime, ctime)
	if err != nil {
		dto.Code = GetErrCode(ErrorTimeFormat)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	if platform == Ios && channel == Ios_Channel {
		plistUrl, err := savePlist(env, app, channel, version, id, title, softwareUrl, fullUrl, displayUrl)
		if err != nil {
			dto.Code = GetErrCode(err)
			dto.Msg = GetMsg(dto.Code)
			c.Data["json"] = dto
			c.ServeJSON()
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
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	err = models.DescListDao.Save(platform, env, app, info.Version, score)
	if err != nil {
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func savePlist(env Environment, app, channel, version, id, title, softwareUrl, fullUrl, displayUrl string) (string, error) {

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
	return getPlistUrl(env, app, version), nil
}
func getPlistUrl(env Environment, app, version string) string {
	var buffer bytes.Buffer
	buffer.WriteString(Url_Start)
	buffer.WriteString(Slash)
	buffer.WriteString(app)
	buffer.WriteString(Slash)
	buffer.WriteString(Plist_Url_Uniform)
	buffer.WriteString(Slash)
	if env == Dev {
		buffer.WriteString(Dev_Str)
	} else {
		buffer.WriteString(Release_Str)
	}
	buffer.WriteString("?version=")
	buffer.WriteString(version)
	return buffer.String()
}
func (c *MainController) Delete() {
	dto := NewSuccessResponseDto()
	residue, _ := c.GetInt("residue", -1)
	if residue < Min_Residue {
		beego.Info("Delete param residue  error !residue:", residue)
		dto.Code = ErrParams
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	var err error
	appParam := c.Ctx.Input.Param(":app")
	app, err := GetApp(appParam)
	if err != nil {
		beego.Info("Delete param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	platformParam := c.Ctx.Input.Param(":platform")
	platform, err := GetPlaform(platformParam)
	if err != nil {
		beego.Info("Delete param platform  error !platform:", platformParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("Delete param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	count, err := models.DescListDao.Count(platform, env, app)
	if err != nil {
		beego.Error("Delete count  error :" + err.Error())
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	if residue < count {
		versions, err := models.DescListDao.Remove(platform, env, app, residue)
		if err != nil {
			dto.Code = GetErrCode(err)
			dto.Msg = GetMsg(dto.Code)
			c.Data["json"] = dto
			c.ServeJSON()
			return
		}
		if versions != nil && len(versions) > 0 {
			models.DescDao.Remove(platform, env, app, versions)
			if platform == Ios {
				err := models.PlistDao.Remove(env, app, versions)
				if err != nil {
					dto.Code = GetErrCode(err)
					dto.Msg = GetMsg(dto.Code)
					c.Data["json"] = dto
					c.ServeJSON()
					return
				}
			}
		}
	}
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
		break
	case Ios:
		item.Platform = Ios_Str
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
	item.Down = info.SoftwareUrl
	item.Time = info.Time
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
