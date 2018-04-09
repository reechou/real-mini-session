package models

type TaskMomentDetail struct {
	TaskMoment  `xorm:"extends" json:"taskMoment"`
	Task        `xorm:"extends" json:"task"`
	SessionInfo `xorm:"extends" json:"user"`
}

func (TaskMomentDetail) TableName() string {
	return "task_moment"
}

func GetTaskMomentDetailList(eventId int64, offset, num int) ([]TaskMomentDetail, error) {
	taskMomentDetailList := make([]TaskMomentDetail, 0)
	err := x.Join("LEFT", "task", "task_moment.task_id = task.id").
		Join("LEFT", "session_info", "task_moment.user_id = session_info.id").
		Where("task_moment.event_id = ?", eventId).Desc("task_moment.created_at").Limit(num, offset).
		Find(&taskMomentDetailList)
	if err != nil {
		return nil, err
	}
	return taskMomentDetailList, nil
}
