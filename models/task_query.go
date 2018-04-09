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

func GetEventTaskTagList(eventId int64) ([]EventTaskTag, error) {
	var taskTags []EventTaskTag
	err := x.Where("event_id = ?", eventId).Find(&taskTags)
	if err != nil {
		return nil, err
	}
	return taskTags, nil
}

func GetTasksFromEventTags(tags []int64) ([]TaskTag, error) {
	var taskTags []TaskTag
	err := x.In("event_task_tag_id", tags).Find(&taskTags)
	if err != nil {
		return nil, err
	}
	return taskTags, nil
}

type TaskTagDetail struct {
	TaskTag      `xorm:"extends" json:"taskTag"`
	EventTaskTag `xorm:"extends" json:"eventTaskTag"`
}

func (TaskTagDetail) TableName() string {
	return "task_tag"
}

func GetTaskTagDetailList(taskId int64) ([]TaskTagDetail, error) {
	tasktagDetailList := make([]TaskTagDetail, 0)
	err := x.Join("LEFT", "event_task_tag", "task_tag.event_task_tag_id = event_task_tag.id").
		Where("task_tag.task_id = ?", taskId).
		Find(&tasktagDetailList)
	if err != nil {
		return nil, err
	}
	return tasktagDetailList, nil
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

type TaskEventDetail struct {
	Task  `xorm:"extends" json:"task"`
	Event `xorm:"extends" json:"event"`
}

func (TaskEventDetail) TableName() string {
	return "task"
}

func GetTaskEventDetailList(userId, start, end int64) ([]TaskEventDetail, error) {
	taskEventList := make([]TaskEventDetail, 0)
	err := x.Join("LEFT", "event", "task.event_id = event.id").
		Where("task.create_user = ? AND task.remind_time >= ? AND task.remind_time <= ?", userId, start, end).
		Find(&taskEventList)
	if err != nil {
		return nil, err
	}
	return taskEventList, nil
}
