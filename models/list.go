package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type List struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	UserId    int64  `xorm:"not null default 0 int index" json:"userId"`
	Name      string `xorm:"not null default '' varchar(128)" json:"name"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateList(info *List) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create list info error: %v", err)
		return err
	}
	holmes.Info("create list info[%v] success.", info)

	return nil
}

func DelList(info *List) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func UpdateList(info *List) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("name", "updated_at").Update(info)
	return err
}

type ListTag struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	UserId    int64  `xorm:"not null default 0 int index" json:"userId"`
	ListId    int64  `xorm:"not null default 0 int index" json:"listId"`
	Name      string `xorm:"not null default '' varchar(128)" json:"name"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateListTag(info *ListTag) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create list tag error: %v", err)
		return err
	}
	holmes.Info("create list tag[%v] success.", info)

	return nil
}

func CreateListTags(list []ListTag) error {
	if len(list) == 0 {
		return nil
	}
	now := time.Now().Unix()
	for i := 0; i < len(list); i++ {
		list[i].CreatedAt = now
		list[i].UpdatedAt = now
	}
	_, err := x.Insert(&list)
	if err != nil {
		holmes.Error("create list tag list error: %v", err)
		return err
	}
	return nil
}

func DelListTag(info *ListTag) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}
