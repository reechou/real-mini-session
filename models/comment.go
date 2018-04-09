package models

import (
	"time"

	"github.com/reechou/holmes"
)

type TaskComment struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	EventId   int64  `xorm:"not null default 0 int index" json:"eventId"`
	UserId    int64  `xorm:"not null default 0 int index" json:"userId"`
	Comment   string `xorm:"not null default '' varchar(512)" json:"comment"`
	CreatedAt int64  `xorm:"not null default 0 int index" json:"createdAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateTaskComment(info *TaskComment) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create task comment info error: %v", err)
		return err
	}
	holmes.Info("create task comment[%v] success.", info)

	return nil
}
