// constants
package constants

import (
	"errors"
)

type ErrCode int32

type Platform int32

type Environment int32

var (
	ErrorFileNotExist     = errors.New("appinhouse: file not exist")
	ErrorUnknown          = errors.New("appinhouse: unknown error")
	ErrorPageOut          = errors.New("appinhouse: page size out of total")
	ErrorParseConf        = errors.New("appinhouse: conf file parse error")
	ErrorParam            = errors.New("appinhouse:param error")
	ErrorCreateFileError  = errors.New("appinhouse:create File fail")
	ErrorIPANotExistError = errors.New("appinhouse:not have IPA ")
	ErrorDeleteFileError  = errors.New("appinhouse:delete File fail")

	errCodeToError = map[string]ErrCode{
		ErrorUnknown.Error():          ErrUnknown,
		ErrorParam.Error():            ErrParams,
		ErrorFileNotExist.Error():     ErrFileNotExist,
		ErrorPageOut.Error():          ErrPageOut,
		ErrorParseConf.Error():        ErrParseConf,
		ErrorCreateFileError.Error():  ErrCreateFileError,
		ErrorIPANotExistError.Error(): ErrIPANotExistError,
		ErrorDeleteFileError.Error():  ErrDeleteFileError,
	}
	Root_Dir    = ""
	Log_Dir     = ""
	Ios_Channel = ""
	Domain      = ""
	platforms   = map[string]Platform{
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
	return "", ErrorUnknown
}

func GetEnvironment(environment string) (Environment, error) {
	if environment == "" {
		return 0, ErrorParam
	}
	if ret, ok := environments[environment]; ok {
		return ret, nil
	}
	return 0, ErrorUnknown
}
func GetPlaform(platform string) (Platform, error) {
	if platform == "" {
		return 0, ErrorParam
	}
	if ret, ok := platforms[platform]; ok {
		return ret, nil
	}
	return 0, ErrorUnknown
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
)
const (
	Dev_Path     = "dev/"
	Release_Path = "release/"
	Android_Path = "android/"
	Ios_Path     = "ios/"

	List_Path   = "list/"
	Data_Path   = "data/"
	Static_Path = "download/"
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
	Http_Str    = "http://"
	Https_Str   = "https://"
	Log_File    = "appinhouse.log"
)
