// constants
package constants

import (
	"errors"
)

type ErrCode int32

var (
	ErrorFileNotExist     = errors.New("appinhouse: file not exist")
	ErrorUnknown          = errors.New("appinhouse: unknown error")
	ErrorPageOut          = errors.New("appinhouse: page size out of total")
	ErrorParseConf        = errors.New("appinhouse: conf file parse error")
	ErrorParam            = errors.New("appinhouse:param error")
	ErrorUserConfigKeyNil = errors.New("appinhouse:user config key is nil")

	errCodeToError = map[string]ErrCode{
		ErrorUnknown.Error():      ErrUnknown,
		ErrorParam.Error():        ErrParams,
		ErrorFileNotExist.Error(): ErrFileNotExist,
		ErrorPageOut.Error():      ErrPageOut,
		ErrorParseConf.Error():    ErrParseConf,
	}
	Root_Dir    = ""
	Ios_Channel = ""
	Domain      = ""
)

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
	ErrUserConfigKeyNil = ErrCode(1005)
)
const (
	App_Dir      = "appinhouse/"
	Dev_Path     = "dev/"
	Release_Path = "release/"
	Android_Path = "android/"
	Ios_Path     = "ios/"

	List_Path                   = "list/"
	Data_Path                   = "data/"
	Android_Release_Static_Path = "/downar"
	Android_Dev_Static_Path     = "/downad"
	Ios_Release_Static_Path     = "/downir"
	Ios_Dev_Static_Path         = "/downid"
)
const (
	Ios       = 1
	Android   = 2
	Dev       = 1
	Release   = 2
	Page_Size = 5
	Max_Page  = 100
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
)
