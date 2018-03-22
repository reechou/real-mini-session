package models

import (
	"time"

	"fmt"
	"github.com/reechou/holmes"
)

type Task struct {
	ID         int64  `xorm:"pk autoincr" json:"id"`
	ListId     int64  `xorm:"not null default 0 int index" json:"listId"`
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

type TaskMember struct {
	ID        int64 `xorm:"pk autoincr"`
	TaskId    int64 `xorm:"not null default 0 int unique(task_member)" json:"taskId"`
	UserId    int64 `xorm:"not null default 0 int unique(task_member)" json:"userId"`
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
