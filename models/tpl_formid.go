package models

import (
	"fmt"
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

func DelTplFormid(info *TplFormid) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func DelExpireTplFormids(userId, expire int64) error {
	_, err := x.Where("user_id = ?", userId).And("expire <= ?", expire).Delete(&TplFormid{})
	if err != nil {
		return err
	}
	return nil
}

func GetTplFormids(userId int64) ([]TplFormid, error) {
	var formids []TplFormid
	err := x.Where("user_id = ?", userId).Find(&formids)
	if err != nil {
		return nil, err
	}
	return formids, nil
}
