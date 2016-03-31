// info.go
package models

import (
	. "appinhouse/inhouse_service/constants"

	"github.com/Unknwon/goconfig"
)

const (
	version     = "version"
	time        = "time"
	description = "description"
	url         = "url"
	channel     = "channel"
	channel_def = "none"
)

type Info struct {
	Version     string
	Time        string
	Description string
	Url         string
	Channel     string
}

func InitConfig(path string) (*Info, error) {
	info := &Info{}
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		return nil, ErrorParseConf
	}
	info.Version, err = cfg.GetValue("", version)
	if err != nil {
		return nil, ErrorParseConf
	}
	info.Time, err = cfg.GetValue("", time)
	if err != nil {
		return nil, ErrorParseConf
	}
	info.Description, err = cfg.GetValue("", description)
	if err != nil {
		return nil, ErrorParseConf
	}
	info.Url, err = cfg.GetValue("", url)
	if err != nil {
		return nil, ErrorParseConf
	}
	info.Channel, err = cfg.GetValue("", channel)
	if err != nil {
		return nil, ErrorParseConf
	}
	return info, nil
}
