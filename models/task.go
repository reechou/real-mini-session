package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type Task struct {
	ID         int64  `xorm:"pk autoincr" json:"id"`
	EventId    int64  `xorm:"not null default 0 int index" json:"eventId"`
	CreateUser int64  `xorm:"not null default 0 int" json:"createUser"`
	Name       string `xorm:"not null default '' varchar(128)" json:"name"`
	Date       string `xorm:"not null default '' varchar(128)" json:"date"`
	Time       string `xorm:"not null default '' varchar(128)" json:"time"`
	RemindTime int64  `xorm:"not null default 0 int index" json:"remindTime"`
	Note       string `xorm:"not null default '' varchar(512)" json:"note"`
	Status     int64  `xorm:"not null default 0 int index" json:"status"`
	CreatedAt  int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt  int64  `xorm:"not null default 0 int" json:"-"`
}

func CreateTask(info *Task) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create task info error: %v", err)
		return err
	}
	holmes.Info("create task info[%v] success.", info)

	return nil
}

func DelTask(info *Task) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(info *Task) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("name", "date", "time", "remind_time", "note", "updated_at").Update(info)
	return err
}

func UpdateTaskStatus(info *Task) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("status", "updated_at").Update(info)
	return err
}

type TaskTag struct {
	ID             int64 `xorm:"pk autoincr" json:"id"`
	TaskId         int64 `xorm:"not null default 0 int index" json:"taskId"`
	EventTaskTagId int64 `xorm:"not null default 0 int index" json:"eventTaskTagId"`
	CreatedAt      int64 `xorm:"not null default 0 int" json:"createdAt"`
}

func CreateTaskTag(info *TaskTag) error {
	now := time.Now().Unix()
	info.CreatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create task tag error: %v", err)
		return err
	}
	holmes.Info("create task tag[%v] success.", info)

	return nil
}

func DelTaskTag(info *TaskTag) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

type EventTaskTag struct {
	ID        int64  `xorm:"pk autoincr" json:"id"`
	EventId   int64  `xorm:"not null default 0 int index" json:"eventId"`
	Name      string `xorm:"not null default '' varchar(128)" json:"name"`
	CreatedAt int64  `xorm:"not null default 0 int" json:"createdAt"`
}

func CreateEventTaskTag(info *EventTaskTag) error {
	now := time.Now().Unix()
	info.CreatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create event task tag error: %v", err)
		return err
	}
	holmes.Info("create event task tag[%v] success.", info)

	return nil
}

func DelEventTaskTag(info *EventTaskTag) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

type TaskMember struct {
	ID        int64 `xorm:"pk autoincr" json:"id"`
	TaskId    int64 `xorm:"not null default 0 int unique(task_member)" json:"taskId"`
	UserId    int64 `xorm:"not null default 0 int unique(task_member)" json:"userId"`
	IfNotify  int64 `xorm:"not null default 0 int" json:"ifNotify"`
	CreatedAt int64 `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateTaskMember(info *TaskMember) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create task member error: %v", err)
		return err
	}
	holmes.Info("create task member[%v] success.", info)

	return nil
}

func CreateTaskMembers(list []TaskMember) error {
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
		holmes.Error("create task member list error: %v", err)
		return err
	}
	return nil
}

func DelTaskMember(info *TaskMember) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}
