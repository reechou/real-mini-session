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

func GetRemindTaskList(start, end int64) ([]Task, error) {
	var tasks []Task
	err := x.Where("status = 0").And("remind_time >= ?", start).And("remind_time <= ?", end).Find(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
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
	TaskMember  `xorm:"extends" json:"taskMember"`
	SessionInfo `xorm:"extends" json:"user"`
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

type TaskInfoDetail struct {
	Task        `xorm:"extends" json:"task"`
	TaskMember  `xorm:"extends" json:"taskMember"`
	SessionInfo `xorm:"extends" json:"user"`
}

func (TaskInfoDetail) TableName() string {
	return "task"
}

func GetTaskInfoDetailList(eventId int64) ([]TaskInfoDetail, error) {
	taskList := make([]TaskInfoDetail, 0)
	err := x.Join("LEFT", "task_member", "task.id = task_member.task_id").
		Join("LEFT", "session_info", "task_member.user_id = session_info.id").
		Where("task.event_id = ?", eventId).
		Find(&taskList)
	if err != nil {
		return nil, err
	}
	return taskList, nil
}
