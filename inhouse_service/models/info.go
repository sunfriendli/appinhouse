// info.go
package models

import (
	. "appinhouse/inhouse_service/constants"

	"bytes"

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

func SaveConfig(info *Info, path string) error {
	c, err := goconfig.LoadFromReader(bytes.NewBuffer([]byte("")))
	if err != nil {
		return ErrorCreateFileError
	}
	c.SetValue("", version, info.Version)
	c.SetValue("", channel, info.Channel)
	c.SetValue("", url, info.Url)
	c.SetValue("", time, info.Time)
	c.SetValue("", description, info.Description)
	err = goconfig.SaveConfigFile(c, path)
	if err != nil {
		return ErrorCreateFileError
	}
	return nil
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
