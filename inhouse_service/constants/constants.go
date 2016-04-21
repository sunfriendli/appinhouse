// constants
package constants

import (
	"errors"
)

type ErrCode int32

type Platform int32

type Environment int32

type ArchiveFileType int32

type PathType int32

var (
	ErrorFileNotExist       = errors.New("appinhouse: file not exist")
	ErrorUnknown            = errors.New("appinhouse: unknown error")
	ErrorPageOut            = errors.New("appinhouse: page size out of total")
	ErrorParseConf          = errors.New("appinhouse: conf file parse error")
	ErrorParam              = errors.New("appinhouse: param error")
	ErrorCreateFileError    = errors.New("appinhouse: create File fail")
	ErrorIPANotExistError   = errors.New("appinhouse: not have IPA ")
	ErrorPlistNotExistError = errors.New("appinhouse: not have Plist ")
	ErrorDeleteFileError    = errors.New("appinhouse: delete File fail")
	ErrorFileExist          = errors.New("appinhouse: file exist")
	errCodeToError          = map[string]ErrCode{
		ErrorUnknown.Error():            ErrUnknown,
		ErrorParam.Error():              ErrParams,
		ErrorFileNotExist.Error():       ErrFileNotExist,
		ErrorPageOut.Error():            ErrPageOut,
		ErrorParseConf.Error():          ErrParseConf,
		ErrorCreateFileError.Error():    ErrCreateFileError,
		ErrorIPANotExistError.Error():   ErrIPANotExistError,
		ErrorDeleteFileError.Error():    ErrDeleteFileError,
		ErrorFileExist.Error():          ErrFileExist,
		ErrorPlistNotExistError.Error(): ErrPlistNotExist,
	}
	errCodeToMsg = map[ErrCode]string{
		ErrOk:               "成功",
		ErrSystemError:      "系统错误",
		ErrUnknown:          "未知错误",
		ErrParams:           "参数错误，请参考API文档",
		ErrFileNotExist:     "文件不存在",
		ErrPageOut:          "page超出最大值",
		ErrParseConf:        "解析描述文件错误",
		ErrCreateFileError:  "创建文件错误",
		ErrIPANotExistError: "归档文件不存在",
		ErrDeleteFileError:  "删除文件错误",
		ErrFileExist:        "文件已存在",
		ErrPlistNotExist:    "plist文件不存在",
	}
	Archive_File_Domain = ""
	Root_Dir            = ""
	Log_Dir             = ""
	Ios_Channel         = ""
	Domain              = ""
	platforms           = map[string]Platform{
		Android_Str: Android,
		Ios_Str:     Ios,
	}
	environments = map[string]Environment{
		Release_Str: Release,
		Dev_Str:     Dev,
	}
	apps            = map[string]string{}
	Full_Size_Image = ""
	Display_Image   = ""
	Min_Residue     = 0
	Page_Size       = 0
	Max_Page        = 0
)

func AddApps(appdirs []string) {
	for _, app := range appdirs {
		if app != "" {
			apps[app] = app
		}
	}
}
func GetApps() []string {
	size := len(apps)
	if size == 0 {
		panic("app is null in app.conf ")
	}
	ret := make([]string, 0, size)
	for _, value := range apps {
		ret = append(ret, value)
	}
	return ret
}
func GetApp(app string) (string, error) {
	if app == "" {
		return "", ErrorParam
	}
	if ret, ok := apps[app]; ok {
		return ret, nil
	}
	return "", ErrorParam
}

func GetEnvironment(environment string) (Environment, error) {
	if environment == "" {
		return 0, ErrorParam
	}
	if ret, ok := environments[environment]; ok {
		return ret, nil
	}
	return 0, ErrorParam
}
func GetPlaform(platform string) (Platform, error) {
	if platform == "" {
		return 0, ErrorParam
	}
	if ret, ok := platforms[platform]; ok {
		return ret, nil
	}
	return 0, ErrorParam
}
func GetErrCode(e error) ErrCode {
	if e == nil {
		return ErrOk
	}
	if err, ok := errCodeToError[e.Error()]; ok {
		return err
	}
	return ErrUnknown
}
func GetMsg(code ErrCode) string {
	if msg, ok := errCodeToMsg[code]; ok {
		return msg
	}
	return errCodeToMsg[ErrUnknown]
}

const (
	ErrOk               = 0
	ErrSystemError      = -1
	ErrUnknown          = ErrCode(-2)
	ErrParams           = ErrCode(1001)
	ErrFileNotExist     = ErrCode(1002)
	ErrPageOut          = ErrCode(1003)
	ErrParseConf        = ErrCode(1004)
	ErrCreateFileError  = ErrCode(1005)
	ErrIPANotExistError = ErrCode(1006)
	ErrDeleteFileError  = ErrCode(1007)
	ErrFileExist        = ErrCode(1008)
	ErrPlistNotExist    = ErrCode(1009)
)
const (
	Dev_Path     = "dev/"
	Release_Path = "release/"
	Android_Path = "android/"
	Ios_Path     = "ios/"

	List_Path         = "list/"
	Data_Path         = "data/"
	Static_Path       = "download/"
	Static_Path_Plist = "plist/"
)
const (
	Ios     = Platform(1)
	Android = Platform(2)
	Dev     = Environment(1)
	Release = Environment(2)
)
const (
	Ios_Str     = "ios"
	Android_Str = "android"
	Release_Str = "release"
	Dev_Str     = "dev"
)
const (
	Plist       = ".plist"
	Apk         = ".apk"
	Ipa         = ".ipa"
	AndroidDown = "androidown"
	IOSDown     = "iosdown"
	Slash       = "/"
	Underline   = "_"
	Http_Str    = "http://"
	Https_Str   = "https://"
	Log_File    = "appinhouse.log"
)
const (
	APK           = ArchiveFileType(0)
	IPA           = ArchiveFileType(1)
	FullSizeImage = ArchiveFileType(2)
	DisplayImage  = ArchiveFileType(3)
)
const (
	List = PathType(0)
	Data = PathType(1)
)
