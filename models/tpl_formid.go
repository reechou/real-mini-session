package models

import (
	"time"

	"github.com/reechou/holmes"
)

type TplFormid struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	UserId    int64  `xorm:"not null default 0 int index" json:"userId"`
	OpenId    string `xorm:"not null default '' varchar(128)" json:"openId"`
	FormId    string `xorm:"not null default '' varchar(128)" json:"formId"`
	Expire    int64  `xorm:"not null default 0 int index" json:"expire"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createdAt"`
}

func CreateTplFormids(list []TplFormid) error {
	if len(list) == 0 {
		return nil
	}
	now := time.Now().Unix()
	for i := 0; i < len(list); i++ {
		list[i].CreatedAt = now
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create form id list error: %v", err)
		return err
	}
	return nil
}
