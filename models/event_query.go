package models

func GetEvent(info *Event) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func GetListEvents(listId int64) ([]Event, error) {
	var listEvents []Event
	err := x.Where("list_id = ?", listId).Find(&listEvents)
	if err != nil {
		return nil, err
	}
	return listEvents, nil
}

type EventMemberDetail struct {
	EventMember `xorm:"extends" json:"eventMember"`
	SessionInfo `xorm:"extends" json:"user"`
}

func (EventMemberDetail) TableName() string {
	return "event_member"
}

func GetEventMemberDetailList(eventId int64) ([]EventMemberDetail, error) {
	eventMemberDetailList := make([]EventMemberDetail, 0)
	err := x.Join("LEFT", "session_info", "event_member.user_id = session_info.id").
		Where("event_member.event_id = ?", eventId).
		Find(&eventMemberDetailList)
	if err != nil {
		return nil, err
	}
	return eventMemberDetailList, nil
}

type ShareEventDetail struct {
	ShareEvent `xorm:"extends" json:"shareEvent"`
	Event      `xorm:"extends" json:"event"`
}

func (ShareEventDetail) TableName() string {
	return "share_event"
}

func GetShareEventDetailList(userId int64) ([]ShareEventDetail, error) {
	shareEventDetailList := make([]ShareEventDetail, 0)
	err := x.Join("LEFT", "event", "share_event.event_id = event.id").
		Where("share_event.user_id = ?", userId).
		Find(&shareEventDetailList)
	if err != nil {
		return nil, err
	}
	return shareEventDetailList, nil
}
