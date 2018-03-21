package models

func GetList(info *List) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func GetListTags(listId int64) ([]ListTag, error) {
	var listTags []ListTag
	err := x.Where("list_id = ?", listId).Find(&listTags)
	if err != nil {
		return nil, err
	}
	return listTags, nil
}

type ListEvent struct {
	List  `xorm:"extends" json:"list"`
	Event `xorm:"extends" json:"event"`
}

func (ListEvent) TableName() string {
	return "list"
}

func GetListEvent(userId int64) ([]ListEvent, error) {
	listEvent := make([]ListEvent, 0)
	err := x.Join("LEFT", "event", "list.id = event.list_id").
		Where("list.user_id = ?", userId).
		Find(&listEvent)
	if err != nil {
		return nil, err
	}
	return listEvent, nil
}
