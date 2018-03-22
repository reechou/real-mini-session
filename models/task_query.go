package models

func GetTask(info *Task) (bool, error) {
	has, err := x.Id(info.ID).Get(info)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func GetTaskMemberList(taskId int64) ([]TaskMember, error) {
	var taskMembers []TaskMember
	err := x.Where("task_id = ?", taskId).Find(&taskMembers)
	if err != nil {
		return nil, err
	}
	return taskMembers, nil
}

type TaskMemberDetail struct {
	TaskMember  `xorm:"extends"`
	SessionInfo `xorm:"extends"`
}

func (TaskMemberDetail) TableName() string {
	return "task_member"
}

func GetTaskMemberDetailList(taskId int64) ([]TaskMemberDetail, error) {
	taskMemberDetailList := make([]TaskMemberDetail, 0)
	err := x.Join("LEFT", "session_info", "task_member.user_id = session_info.id").
		Where("task_member.task_id = ?", taskId).
		Find(&taskMemberDetailList)
	if err != nil {
		return nil, err
	}
	return taskMemberDetailList, nil
}
