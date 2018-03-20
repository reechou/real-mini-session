package models

type ListEvent struct {
	List  `xorm:"extends"`
	Event `xorm:"extends"`
}

func (ListEvent) TableName() string {
	return "list"
}

func GetListEvent(userId int64) ([]ListEvent, error) {
	listEvent := make([]ListEvent, 0)
	err := x.Join("LEFT", "event", "list.id = event.list_id").
		Where("task.user_id = ?", userId).
		Find(&listEvent)
	if err != nil {
		return nil, err
	}
	return listEvent, nil
}
