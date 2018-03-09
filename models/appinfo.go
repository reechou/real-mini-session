package models

import (
	"github.com/reechou/holmes"
)

type AppInfo struct {
	ID              int64  `xorm:"pk autoincr"`
	AppId           string `xorm:"not null default '' varchar(128) unique" json:"appid"`
	Secret          string `xorm:"not null default '' varchar(300)" json:"secret"`
	LoginDuration   int64  `xorm:"not null default 30 int" json:"login_duration"`
	SessionDuration int64  `xorm:"not null default 2592000 int" json:"session_duration"`
	CreatedAt       int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt       int64  `xorm:"not null default 0 int" json:"-"`
}

func GetAppInfo(info *AppInfo) (bool, error) {
	has, err := x.Where("app_id = ?", info.AppId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find appinfo from [%v]", info)
		return false, nil
	}
	return true, nil
}
