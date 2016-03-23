package controllers

import (
	"io/ioutil"
	//	"path/filepath"

	"os"

	"sort"

	"appinhouse/models"

	"appinhouse/util"

	. "appinhouse/constants"

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

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	userAgent := c.Ctx.Request.UserAgent()
	osType := util.CheckAgent(userAgent)
	switch osType {
	case util.OS_ANDROID:
		c.TplName = "index_mobile.html"
		break
	case util.OS_IOS:
		c.TplName = "index_mobile.html"
		break
	default:
		c.TplName = "index_mobile.html"
		break
	}

}

func (c *MainController) Last() {
	dto := NewSuccessItemsResponseDto()
	userAgent := c.Ctx.Request.UserAgent()
	osType := util.CheckAgent(userAgent)
	ret := make([]*ItemDto, 0, 4)
	switch osType {
	case util.OS_ANDROID:
		androiddev, err := getLastOne(Android, Dev)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		androidre, err := getLastOne(Android, Release)
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
		iosdev, err := getLastOne(Ios, Dev)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release)
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
		iosdev, err := getLastOne(Ios, Dev)
		if err != nil {
			dto.Code = GetErrCode(err)
			break
		}
		iosre, err := getLastOne(Ios, Release)
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
		//		androiddev, err := getLastOne(Android, Dev)
		//		if err != nil {
		//			dto.Code = GetErrCode(err)
		//			break
		//		}
		//		androidre, err := getLastOne(Android, Release)
		//		if err != nil {
		//			dto.Code = GetErrCode(err)
		//			break
		//		}
		//		iosdev, _ := getLastOne(Ios, Dev)
		//		iosre, _ := getLastOne(Ios, Release)
		//		if androiddev != nil {
		//			ret = append(ret, androiddev)
		//		}
		//		if androidre != nil {
		//			ret = append(ret, androidre)
		//		}

		//		if iosdev != nil {
		//			ret = append(ret, iosdev)
		//		}
		//		if iosre != nil {
		//			ret = append(ret, iosre)
		//		}
		break
	}

	dto.Items = ret
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) DevList() {
	dto := NewSuccessItemsResponsePageDto()
	page, _ := c.GetInt("page", 1)
	if page <= 0 || page > Max_Page {
		dto.Code = ErrParams
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
	var err error
	switch osType {
	case util.OS_ANDROID:
		ret, total, err = getList(Android, Dev, start, end)
		break
	case util.OS_IOS:
		ret, total, err = getList(Ios, Dev, start, end)
		break
	default:
		ret, total, err = getList(Android, Dev, start, end)
		break
	}

	if err != nil {
		dto.Code = GetErrCode(err)
	} else {
		dto.Items = ret
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(total, Page_Size)
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) ReleaseList() {
	dto := NewSuccessItemsResponsePageDto()
	page, _ := c.GetInt("page", 1)
	if page <= 0 || page > Max_Page {
		dto.Code = ErrParams
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)
	userAgent := c.Ctx.Request.UserAgent()
	osType := util.CheckAgent(userAgent)
	var err error
	var ret *ItemsDto
	total := 0
	switch osType {
	case util.OS_ANDROID:
		ret, total, err = getList(Android, Release, start, end)

		break
	case util.OS_IOS:
		ret, total, err = getList(Ios, Release, start, end)
		break
	default:
		ret, total, err = getList(Android, Release, start, end)
		break
	}
	if err != nil {
		dto.Code = GetErrCode(err)
	} else {
		dto.Items = ret
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(total, Page_Size)
	}
	c.Data["json"] = dto
	c.ServeJSON()
}
func (c *MainController) AndroidDevList() {
	c.getList(Android, Dev)
}
func (c *MainController) AndroidReleaseList() {
	c.getList(Android, Release)
}
func (c *MainController) IosDevList() {
	c.getList(Ios, Dev)
}
func (c *MainController) IosReleaseList() {
	c.getList(Ios, Release)
}
func (c *MainController) getList(platform, env int) {
	dto := NewSuccessItemsResponsePageDto()
	page, _ := c.GetInt("page", 1)
	if page <= 0 || page > Max_Page {
		dto.Code = ErrParams
		c.Data["json"] = dto
		c.ServeJSON()
		return
	}
	start := util.GetStartForPage(page, Page_Size)
	end := util.GetEndForPage(page, Page_Size)

	ret, total, err := getList(platform, env, start, end)
	if err != nil {
		dto.Code = GetErrCode(err)
	} else {
		dto.Items = ret
		dto.Page = page
		dto.TotalPage = util.GetTotalPage(total, Page_Size)
	}
	c.Data["json"] = dto
	c.ServeJSON()
}

func getLastOne(platform int, environment int) (*ItemDto, error) {
	repositorys, err := getAll(platform, environment)
	if err != nil {
		beego.Error("getLastOne  call getall error:" + err.Error())
		return nil, err
	}
	if len(*repositorys) == 0 {
		return nil, err
	}
	a := *repositorys
	path := a[0].Name
	info, err := models.InitConfig(getListPath(platform, environment) + path)
	if err != nil {
		beego.Error("getLastOne  call InitConfig:" + err.Error())
		return nil, err
	}
	return converInfoTOItem(info, platform, environment), nil
}
func converInfoTOItem(info *models.Info, platform int, environment int) *ItemDto {

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
	item.Down = getDownPath(platform, environment, info.Channel, info.Version)
	item.Time = info.Time
	item.Url = info.Description
	item.Version = info.Version

	return item
}
func getStaticPath(platform, environment int) string {
	ret := ""
	if platform == Ios {
		if environment == Dev {
			ret = Ios_Dev_Static_Path
		} else {
			ret = Ios_Release_Static_Path
		}

	} else {
		if environment == Dev {
			ret = Android_Dev_Static_Path

		} else {
			ret = Android_Release_Static_Path
		}
	}
	return ret
}
func getDownPath(platform, environment int, channel, version string) string {
	data := getDataPath(platform, environment) + version + "/"
	files, err := ioutil.ReadDir(data)
	staticPath := getStaticPath(platform, environment)
	staticPath = staticPath + "/" + version + "/"
	if err != nil {
		return ""
	}
	if platform == Ios {
		if channel == Ios_Channel {
			for _, file := range files {
				if strings.HasSuffix(file.Name(), Plist) {
					return Domain + staticPath + file.Name()
				}
			}
		} else {
			for _, file := range files {

				if strings.HasSuffix(file.Name(), Ipa) {
					return staticPath + file.Name()
				}
			}
		}

	} else {
		for _, file := range files {
			if strings.HasSuffix(file.Name(), Apk) {
				return staticPath + file.Name()
			}
		}
	}
	return ""
}

func getList(platform, environment, start, end int) (*ItemsDto, int, error) {

	repositorys, err := getAll(platform, environment)
	if err != nil {
		beego.Error("getList call  getAll error :" + err.Error())
		return nil, 0, err
	}

	res := *repositorys
	size := len(*repositorys)
	if start > size {
		beego.Error("getList call page start greate than size" + err.Error())
		return nil, 0, ErrorPageOut
	}
	if end > size {
		end = size - 1
	}

	items := new(ItemsDto)
	for i := start; i <= end; i++ {
		info, err := models.InitConfig(getListPath(platform, environment) + res[i].Name)
		if err != nil {
			beego.Error("getList call  models.InitConfig error:" + err.Error())
			return nil, 0, ErrorParseConf
		}
		*items = append(*items, converInfoTOItem(info, platform, environment))
	}

	return items, size, nil
}
func getAll(platform int, environment int) (*FileRepos, error) {
	file := getListPath(platform, environment)
	files, _ := ioutil.ReadDir(file)
	repositorys, _ := sortRepository(files)
	return repositorys, nil
}
func InitDirectory() {
	createFile(getListPath(Ios, Dev))
	createFile(getListPath(Android, Dev))
	createFile(getListPath(Ios, Release))
	createFile(getListPath(Android, Release))
	createFile(getDataPath(Ios, Dev))
	createFile(getDataPath(Android, Dev))
	createFile(getDataPath(Ios, Release))
	createFile(getDataPath(Android, Release))
}
func GetDownDirectory(platform int, environment int) string {
	return getDataPath(platform, environment)
}

func getListPath(platform int, environment int) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Root_Dir + App_Dir + Dev_Path + Ios_Path + List_Path
		} else {
			file = Root_Dir + App_Dir + Release_Path + Ios_Path + List_Path
		}

	} else {
		if environment == Dev {
			file = Root_Dir + App_Dir + Dev_Path + Android_Path + List_Path
		} else {
			file = Root_Dir + App_Dir + Release_Path + Android_Path + List_Path
		}
	}
	return file
}
func getDataPath(platform int, environment int) string {
	file := ""
	if platform == Ios {
		if environment == Dev {
			file = Root_Dir + App_Dir + Dev_Path + Ios_Path + Data_Path
		} else {
			file = Root_Dir + App_Dir + Release_Path + Ios_Path + Data_Path
		}

	} else {
		if environment == Dev {
			file = Root_Dir + App_Dir + Dev_Path + Android_Path + Data_Path
		} else {
			file = Root_Dir + App_Dir + Release_Path + Android_Path + Data_Path
		}
	}
	return file
}
func createFile(dir string) (string, error) {
	src := dir
	if isExist(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0777); err != nil {
		if os.IsPermission(err) {
			beego.Error("permission denied path:" + src)
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

func unlink(file string) error {
	return os.Remove(file)
}
