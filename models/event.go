package models

import (
	"fmt"
	"time"

	"github.com/reechou/holmes"
)

type Event struct {
	ID          int64  `xorm:"pk autoincr" json:"id"`
	ListId      int64  `xorm:"not null default 0 int index" json:"listId"`
	OwnerUserId int64  `xorm:"not null default 0 int index" json:"ownerUserId"`
	Name        string `xorm:"not null default '' varchar(128)" json:"name"`
	CreatedAt   int64  `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt   int64  `xorm:"not null default 0 int" json:"-"`

	TaskNum int64 `json:"taskNum"`
}

func CreateEvent(info *Event) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create event info error: %v", err)
		return err
	}
	holmes.Info("create event info[%v] success.", info)

	return nil
}

func DelEvent(info *Event) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvent(info *Event) error {
	info.UpdatedAt = time.Now().Unix()
	_, err := x.ID(info.ID).Cols("name", "updated_at").Update(info)
	return err
}

type EventMember struct {
	ID        int64 `xorm:"pk autoincr" json:"id"`
	EventId   int64 `xorm:"not null default 0 int unique(event_member)" json:"eventId"`
	UserId    int64 `xorm:"not null default 0 int unique(event_member)" json:"userId"`
	CreatedAt int64 `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateEventMember(info *EventMember) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create event member error: %v", err)
		return err
	}
	holmes.Info("create event member[%v] success.", info)

	return nil
}

func DelEventMember(info *EventMember) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func DelEventMemberFromUserEvent(info *EventMember) error {
	if info.EventId == 0 || info.UserId == 0 {
		return fmt.Errorf("del userid or eventid cannot be nil.")
	}
	_, err := x.Where("event_id = ?", info.EventId).And("user_id = ?", info.UserId).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

type ShareEvent struct {
	ID        int64 `xorm:"pk autoincr" json:"id"`
	UserId    int64 `xorm:"not null default 0 int unique(user_share_event)" json:"userId"`
	EventId   int64 `xorm:"not null default 0 int unique(user_share_event)" json:"eventId"`
	CreatedAt int64 `xorm:"not null default 0 int" json:"createdAt"`
	UpdatedAt int64 `xorm:"not null default 0 int" json:"-"`
}

func CreateShareEvent(info *ShareEvent) error {
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now

	_, err := x.Insert(info)
	if err != nil {
		holmes.Error("create share event error: %v", err)
		return err
	}
	holmes.Info("create share event[%v] success.", info)

	return nil
}

func DelShareEvent(info *ShareEvent) error {
	if info.ID == 0 {
		return fmt.Errorf("del id cannot be nil.")
	}
	_, err := x.ID(info.ID).Delete(info)
	if err != nil {
		return err
	}
	return nil
}

func DelShareEventFromUserEvent(info *ShareEvent) error {
	if info.EventId == 0 || info.UserId == 0 {
		return fmt.Errorf("del userid or eventid cannot be nil.")
	}
	_, err := x.Where("user_id = ?", info.UserId).And("event_id = ?", info.EventId).Delete(info)
	if err != nil {
		return err
	}
	return nil
}
