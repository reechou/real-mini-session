package models

type TaskCommentDetail struct {
	TaskComment `xorm:"extends" json:"taskComment"`
	SessionInfo `xorm:"extends" json:"user"`
}

func (TaskCommentDetail) TableName() string {
	return "task_comment"
}

func GetTaskCommentDetailList(eventId int64, offset, num int) ([]TaskCommentDetail, error) {
	taskCommentDetailList := make([]TaskCommentDetail, 0)
	err := x.Join("LEFT", "session_info", "task_comment.user_id = session_info.id").
		Where("task_comment.event_id = ?", eventId).Decr("task_comment.created_at").Limit(num, offset).
		Find(&taskCommentDetailList)
	if err != nil {
		return nil, err
	}
	return taskCommentDetailList, nil
}
