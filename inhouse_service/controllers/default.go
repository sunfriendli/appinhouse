package controllers

import (
	"io/ioutil"

	"os"

	"sort"

	"appinhouse/inhouse_service/models"

	"appinhouse/inhouse_service/util"

	. "appinhouse/inhouse_service/constants"

	"strings"

	"github.com/astaxie/beego"
)

type FileRepos []*Repository

type Repository struct {
	Name     string
	FileTime int64
}

func (r FileRepos) Len() int {
	return len(r)
}

func (r FileRepos) Less(i, j int) bool {
	return r[i].FileTime > r[j].FileTime
}

func (r FileRepos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func InitDirectory() {
	apps := GetApps()
	for _, app := range apps {
		createFile(getListPath(Ios, Dev, app))
		createFile(getListPath(Android, Dev, app))
		createFile(getListPath(Ios, Release, app))
		createFile(getListPath(Android, Release, app))
		createFile(getDataPath(Ios, Dev, app))
		createFile(getDataPath(Android, Dev, app))
		createFile(getDataPath(Ios, Release, app))
		createFile(getDataPath(Android, Release, app))
	}
}
func CheckFtpDirectory() {
	apps := GetApps()
	for _, app := range apps {
		exist := isExist(getFtpDataPath(Ios, Dev, app))
		if !exist {
			panic(app + " directory is error. path:" + getFtpDataPath(Ios, Dev, app))
		}
		exist = isExist(getFtpDataPath(Android, Dev, app))
		if !exist {
			panic(app + " directory is error. path:" + getFtpDataPath(Android, Dev, app))
		}
		exist = isExist(getFtpDataPath(Ios, Release, app))
		if !exist {
			panic(app + " directory is error. path:" + getFtpDataPath(Ios, Release, app))
		}
		exist = isExist(getFtpDataPath(Android, Release, app))
		if !exist {
			panic(app + " directory is error. path:" + getFtpDataPath(Android, Release, app))
		}
	}
}
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
	var ret *ItemsDto
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
	id := c.GetString("id")
	version := c.GetString("version")
	title := c.GetString("title")
	if id == "" || version == "" || title == "" {
		beego.Info("PList Params fail .id:", id, "version:", version, "title:", title)
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
		beego.Info("PList param app  error !app:", appParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	envParam := c.Ctx.Input.Param(":environment")
	env, err := GetEnvironment(envParam)
	if err != nil {
		beego.Info("PList param environment  error !environment:", envParam)
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	ftpData := getFtpDataPath(Ios, env, app) + version + Slash
	ftpFiles, err := ioutil.ReadDir(ftpData)
	if err != nil {
		beego.Info("PList error.read ftp dir:", ftpData)
		dto.Code = ErrIPANotExistError
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	} else {
		if len(ftpFiles) == 0 {
			beego.Info("PList files len is zero.read dir:", ftpData)
			dto.Code = ErrIPANotExistError
			dto.Msg = GetMsg(dto.Code)
			c.Data["json"] = dto
			c.ServeJSON()
			return
		}
	}
	file := getDataPath(Ios, env, app) + version
	if !isExist(file) {
		createFile(file)
	}
	plistName := file + Slash + version + Plist
	if isExist(plistName) {
		beego.Info("PList is exist!path:", plistName)
		dto.Code = ErrFileExist
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	fullUrl := ""
	softwareUrl := ""
	displayUrl := ""
	staticPath := Domain + getStaticPath(Ios, env, app) + version + Slash
	for _, file := range ftpFiles {
		if strings.HasSuffix(file.Name(), Ipa) {
			softwareUrl = Https_Str + staticPath + file.Name()
		}
		if file.Name() == Full_Size_Image {
			fullUrl = Http_Str + staticPath + file.Name()
		}
		if file.Name() == Display_Image {
			displayUrl = Http_Str + staticPath + file.Name()
		}
	}

	plist, err := models.NewPlist(id, version, title, softwareUrl, fullUrl, displayUrl)

	if err != nil {
		beego.Info("PList NewPlist error :", err.Error())
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	_, err = filePutContent(plistName, plist)
	if err != nil {
		beego.Info("PList CreatPlist fail :", err.Error())
		dto.Code = ErrCreateFileError
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) Desc() {
	dto := NewSuccessResponseDto()

	version := c.GetString("version")
	ctime := c.GetString("time")
	description := c.GetString("description")
	url := c.GetString("url")
	channel := c.GetString("channel")
	if ctime == "" || version == "" || url == "" || channel == "" {
		beego.Info("Desc Params fail version:", version, "time:", ctime, "description:", description, "url:", url, "channel:", channel)
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

	file := getListPath(platform, env, app) + channel + "_" + version + ".txt"
	exist := isExist(file)
	if exist {
		dto.Code = ErrFileExist
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	ftpData := getFtpDataPath(platform, env, app) + version
	ftpFiles, err := ioutil.ReadDir(ftpData)
	if err != nil {
		beego.Info("Desc error.read ftp dir:", ftpData)
		dto.Code = ErrIPANotExistError
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	} else {
		if len(ftpFiles) == 0 {
			beego.Info("ftpData files len is zero.read dir:", ftpData)
			dto.Code = ErrIPANotExistError
			dto.Msg = GetMsg(dto.Code)
			c.Data["json"] = dto
			c.ServeJSON()
			return
		}
	}
	hasFile := false
	for _, file := range ftpFiles {
		if platform == Ios {
			if strings.HasSuffix(file.Name(), Ipa) {
				hasFile = true
				break
			}
		} else {
			if strings.HasSuffix(file.Name(), Apk) {
				hasFile = true
				break
			}
		}
	}
	if !hasFile {
		if len(ftpFiles) == 0 {
			beego.Info("ftpData files not have ipa or apk.read dir:", ftpData)
			dto.Code = ErrIPANotExistError
			dto.Msg = GetMsg(dto.Code)
			c.Data["json"] = dto
			c.ServeJSON()
			return
		}
	}
	info := new(models.Info)
	info.Channel = channel
	info.Description = description
	info.Version = version
	info.Time = ctime
	info.Url = url
	err = models.SaveConfig(info, file)
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
	repositorys, err := getAll(platform, env, app)
	if err != nil {
		beego.Info("Delete call  getAll error :" + err.Error())
		dto.Code = GetErrCode(err)
		dto.Msg = GetMsg(dto.Code)
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}

	res := *repositorys
	size := len(res)

	if residue < size {
		for i := residue; i < size; i++ {
			list := getListPath(platform, env, app) + res[i].Name
			info, err := models.InitConfig(list)
			if err != nil {
				beego.Info("Delete call  models.InitConfig error:"+err.Error(), "path:", list)
				dto.Code = GetErrCode(err)
				c.Data["json"] = dto
				dto.Msg = GetMsg(dto.Code)
				c.ServeJSON()
				return
			}
			data := getDataPath(platform, env, app) + info.Version
			err = removeAll(data)
			if err != nil {
				beego.Info("Delete call  remove error:"+err.Error(), "path:", data)
				dto.Code = GetErrCode(err)
				dto.Msg = GetMsg(dto.Code)
				c.Data["json"] = dto
				c.ServeJSON()
				return
			}
			err = remove(list)
			if err != nil {
				beego.Info("Delete call remove error:"+err.Error(), "path:", list)
				dto.Code = GetErrCode(err)
				dto.Msg = GetMsg(dto.Code)
				c.Data["json"] = dto
				c.ServeJSON()
				return
			}

		}
	}
	c.Data["json"] = dto
	c.ServeJSON()
}

func getLastOne(platform Platform, environment Environment, app string) (*ItemDto, error) {
	repositorys, err := getAll(platform, environment, app)
	if err != nil {
		beego.Error("getLastOne  call getall error:" + err.Error())
		return nil, err
	}
	if len(*repositorys) == 0 {
		return nil, nil
	}
	a := *repositorys
	path := a[0].Name
	info, err := models.InitConfig(getListPath(platform, environment, app) + path)
	if err != nil {
		beego.Error("getLastOne  call InitConfig:" + err.Error())
		return nil, err
	}
	return converInfoTOItem(info, platform, environment, app), nil
}
func converInfoTOItem(info *models.Info, platform Platform, environment Environment, app string) *ItemDto {

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
	item.Down = getDownPath(platform, environment, info.Channel, info.Version, app)
	item.Time = info.Time
	item.Url = info.Url
	item.Version = info.Version

	return item
}
func getStaticPath(platform Platform, environment Environment, app string) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Slash + Static_Path + app + Slash + Dev_Path + Ios_Path + Data_Path
		} else {
			file = Slash + Static_Path + app + Slash + Release_Path + Ios_Path + Data_Path
		}

	} else {
		if environment == Dev {
			file = Slash + Static_Path + app + Slash + Dev_Path + Android_Path + Data_Path
		} else {
			file = Slash + Static_Path + app + Slash + Release_Path + Android_Path + Data_Path
		}
	}
	return file
}
func getDownPath(platform Platform, environment Environment, channel, version, app string) string {
	data := getDataPath(platform, environment, app) + version + Slash
	ftpData := getFtpDataPath(platform, environment, app) + version + Slash
	files, err := ioutil.ReadDir(data)
	if err != nil {
		beego.Info("getDownPath error.read dir:", data)
		return ""
	} else {
		if len(files) == 0 {
			beego.Info("getDownPath files len is zero.read dir:", data)
			return ""
		}
	}
	ftpFiles, err := ioutil.ReadDir(ftpData)
	if err != nil {
		beego.Info("getDownPath error.read ftp dir:", ftpData)
		return ""
	} else {
		if len(ftpFiles) == 0 {
			beego.Info("getDownPath files len is zero.read dir:", ftpData)
			return ""
		}
	}
	staticPath := ""
	staticPath = getStaticPath(platform, environment, app) + version + Slash
	if err != nil {
		return ""
	}
	if platform == Ios {
		if channel == Ios_Channel {
			for _, file := range files {
				if strings.HasSuffix(file.Name(), Plist) {
					return Https_Str + Domain + staticPath + file.Name()
				}
			}
		} else {
			for _, file := range ftpFiles {

				if strings.HasSuffix(file.Name(), Ipa) {
					return staticPath + file.Name()
				}
			}
		}

	} else {
		for _, file := range ftpFiles {
			if strings.HasSuffix(file.Name(), Apk) {
				return staticPath + file.Name()
			}
		}
	}
	return ""
}

func getList(platform Platform, environment Environment, start, end int, app string) (*ItemsDto, int, error) {

	repositorys, err := getAll(platform, environment, app)
	if err != nil {
		beego.Error("getList call  getAll error :" + err.Error())
		return nil, 0, err
	}

	res := *repositorys
	size := len(*repositorys)
	if start > size {
		beego.Info("getList call page start greate than size" + err.Error())
		return nil, 0, ErrorPageOut
	}
	if end > size {
		end = size - 1
	}

	items := new(ItemsDto)
	for i := start; i <= end; i++ {
		info, err := models.InitConfig(getListPath(platform, environment, app) + res[i].Name)
		if err != nil {
			beego.Info("getList call  models.InitConfig error:" + err.Error())
			return nil, 0, ErrorParseConf
		}
		*items = append(*items, converInfoTOItem(info, platform, environment, app))
	}

	return items, size, nil
}
func getAll(platform Platform, environment Environment, app string) (*FileRepos, error) {
	file := getListPath(platform, environment, app)
	files, _ := ioutil.ReadDir(file)
	repositorys, _ := sortRepository(files)
	return repositorys, nil
}
func getListPath(platform Platform, environment Environment, app string) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Root_Dir + app + Slash + Dev_Path + Ios_Path + List_Path
		} else {
			file = Root_Dir + app + Slash + Release_Path + Ios_Path + List_Path
		}

	} else {
		if environment == Dev {
			file = Root_Dir + app + Slash + Dev_Path + Android_Path + List_Path
		} else {
			file = Root_Dir + app + Slash + Release_Path + Android_Path + List_Path
		}
	}
	return file
}
func getDataPath(platform Platform, environment Environment, app string) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Root_Dir + app + Slash + Dev_Path + Ios_Path + Data_Path
		} else {
			file = Root_Dir + app + Slash + Release_Path + Ios_Path + Data_Path
		}

	} else {
		if environment == Dev {
			file = Root_Dir + app + Slash + Dev_Path + Android_Path + Data_Path
		} else {
			file = Root_Dir + app + Slash + Release_Path + Android_Path + Data_Path
		}
	}
	return file
}
func getFtpDataPath(platform Platform, environment Environment, app string) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Ftp_Root_Dir + app + Slash + Dev_Path + Ios_Path + Data_Path
		} else {
			file = Ftp_Root_Dir + app + Slash + Release_Path + Ios_Path + Data_Path
		}

	} else {
		if environment == Dev {
			file = Ftp_Root_Dir + app + Slash + Dev_Path + Android_Path + Data_Path
		} else {
			file = Ftp_Root_Dir + app + Slash + Release_Path + Android_Path + Data_Path
		}
	}
	return file
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
func sortRepository(files []os.FileInfo) (*FileRepos, error) {
	if len(files) <= 0 {
		return new(FileRepos), nil
	}

	result := new(FileRepos)

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			*result = append(*result, &Repository{Name: file.Name(), FileTime: file.ModTime().Unix()})
		}
	}

	sort.Sort(result)

	return result, nil
}

func remove(file string) error {
	err := os.Remove(file)
	if err != nil {
		return ErrorDeleteFileError
	}
	return err
}
func removeAll(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return ErrorDeleteFileError
	}
	return err
}
func filePutContent(file string, content []byte) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.Write(content)
}
