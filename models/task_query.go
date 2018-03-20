package models

func GetTaskMemberList(taskId int64) ([]TaskMember, error) {
	var taskMembers []TaskMember
	err := x.Where("task_id = ?", taskId).Find(&taskMembers)
	if err != nil {
		return nil, err
	}
	return taskMembers, nil
}

type TaskDetail struct {
	Task        `xorm:"extends"`
	TaskMember  `xorm:"extends"`
	SessionInfo `xorm:"extends"`
}

func (TaskDetail) TableName() string {
	return "task"
}

func GetTaskDetail(taskId int64) ([]TaskDetail, error) {
	taskDetail := make([]TaskDetail, 0)
	err := x.Join("LEFT", "task_member", "task.id = task_member.task_id").
		Join("LEFT", "session_info", "task_member.user_id = session_info.id").
		Where("task.id = ?", taskId).
		Find(&taskDetail)
	if err != nil {
		return nil, err
	}
	return taskDetail, nil
}
