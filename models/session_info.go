package models

import (
	"time"

	"github.com/reechou/holmes"
)

type SessionInfo struct {
	ID         int64  `xorm:"pk autoincr"`
	AppId      int64  `xorm:"not null default 0 int unique(user)" json:"-"`
	OpenId     string `xorm:"not null default '' varchar(128) unique(user)" json:"openId"`
	UnionId    string `xorm:"not null default '' varchar(128)" json:"unionId"`
	NickName   string `xorm:"not null default '' varchar(256)" json:"nickName"`
	AvatarUrl  string `xorm:"not null default '' varchar(512)" json:"avatarUrl"`
	Gender     int64  `xorm:"not null default 0 int" json:"gender"`
	Country    string `xorm:"not null default '' varchar(128)" json:"country"`
	Province   string `xorm:"not null default '' varchar(128)" json:"province"`
	City       string `xorm:"not null default '' varchar(128)" json:"city"`
	Skey       string `xorm:"not null default '' varchar(128)" json:"-"`
	SessionKey string `xorm:"not null default '' varchar(128)" json:"-"`
	CreatedAt  int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt  int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateSessionInfo(info *SessionInfo) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create session info error: %v", err)
		return err
	}
	holmes.Info("create session info[%v] success.", info)

	return nil
}

func GetSessionInfo(info *SessionInfo) (bool, error) {
	has, err := x.Where("app_id = ?", info.AppId).And("open_id = ?", info.OpenId).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find session info from [%v]", info)
		return false, nil
	}
	return true, nil
}

func GetSessionInfoFromId(info *SessionInfo) (bool, error) {
	has, err := x.ID(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		holmes.Debug("cannot find session info from [%v]", info)
		return false, nil
	}
	return true, nil
}
