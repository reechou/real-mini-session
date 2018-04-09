package models

import (
	"time"

	"github.com/reechou/holmes"
)

type TaskMoment struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	EventId   int64  `xorm:"not null default 0 int index" json:"evendId"`
	TaskId    int64  `xorm:"not null default 0 int index" json:"taskId"`
	UserId    int64  `xorm:"not null default 0 int index" json:"userId"`
	OprType   int64  `xorm:"not null default 0 int index" json:"oprType"`
	OprInfo   string `xorm:"not null default '' varchar(128)" json:"oprInfo"`
	Detail    string `xorm:"not null default '' varchar(512)" json:"detail"`
	CreatedAt int64  `xorm:"not null default 0 int index" json:"createdAt"`
	UpdatedAt int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateTaskMoment(info *TaskMoment) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create task moment info error: %v", err)
		return err
	}
	holmes.Info("create task moment[%v] success.", info)

	return nil
}
