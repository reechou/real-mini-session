package models

func GetListEvents(listId int64) ([]Event, error) {
	var listEvents []Event
	err := x.Where("list_id = ?", listId).Find(&listEvents)
	if err != nil {
		return nil, err
	}
	return listEvents, nil
}
