package server

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/config"
	"github.com/reechou/real-mini-session/models"
	"github.com/reechou/real-mini-session/wechat"
	"github.com/robfig/cron"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
)

const (
	REMIND_TASK_PAGE  = "pages/events/task/task?eventid=%d"
	RECEIVE_TASK_PAGE = "pages/events/task/task?eventid=%d&tag=NEW&taskid=%d"
	DONE_TASK_PAGE    = "pages/events/task/task?eventid=%d&tag=DONE&taskid=%d"
)

type ListReminder struct {
	cfg *config.Config
	w   *Wechat

	stop chan struct{}
}

func NewListReminder(cfg *config.Config, w *Wechat) *ListReminder {
	lr := &ListReminder{
		cfg: cfg,
		w:   w,
	}
	if lr.cfg.ListRemindCron == "" {
		holmes.Error("list remind cron time is nil, maybe list remind not used.")
		return lr
	}

	go lr.Run()

	return lr
}

func (lr *ListReminder) Stop() {
	close(lr.stop)
}

func (lr *ListReminder) TaskDone(t *models.Task) {
	formid, err := lr.getFormId(t.CreateUser)
	if err != nil {
		holmes.Error("task done tpl msg get form id error: %v", err)
		return
	}
	if formid == nil {
		holmes.Debug("task create user[%d] don't have form id", t.CreateUser)
		return
	}
	tplMsg := &wechat.WechatTplMsg{
		ToUser: formid.OpenId,
		TplId:  lr.cfg.DoneTaskTplId,
		Page:   fmt.Sprintf(DONE_TASK_PAGE, t.EventId, t.ID),
		FormId: formid.FormId,
		Data: &wechat.TaskDoneTplMsg{
			Keyword1: &template.DataItem{
				Value: t.Name,
				Color: "#008B8B",
			},
			Keyword2: &template.DataItem{
				Value: time.Now().Format("2006-01-02 15:04"),
				Color: "#008B8B",
			},
			Keyword3: &template.DataItem{
				Value: "已完成任务",
				Color: "#008B8B",
			},
		},
	}
	lr.w.Send(tplMsg)
}

func (lr *ListReminder) TaskReceive(t *models.Task, userId int64) {
	formid, err := lr.getFormId(userId)
	if err != nil {
		holmes.Error("task done tpl msg get form id error: %v", err)
		return
	}
	if formid == nil {
		holmes.Debug("task create user[%d] don't have form id", t.CreateUser)
		return
	}
	members, err := models.GetTaskMemberDetailList(t.ID)
	if err != nil {
		holmes.Error("get task member detail list error: %v", err)
		return
	}
	taskMembers := ""
	for i := 0; i < len(members); i++ {
		if i == 3 {
			taskMembers += " ..."
			break
		}
		taskMembers += fmt.Sprintf(" %s", members[i].SessionInfo.NickName)
	}
	createUser := &models.SessionInfo{ID: t.CreateUser}
	has, err := models.GetSessionInfoFromId(createUser)
	if err != nil {
		holmes.Error("get create user error: %v", err)
		return
	}
	if !has {
		createUser.NickName = "私人助手"
	}
	remindTime := "暂未定截止时间"
	if t.RemindTime != 0 {
		remindTime = time.Unix(t.RemindTime, 0).Format("2006.01.02 15:04") + " 截止"
	}
	tplMsg := &wechat.WechatTplMsg{
		ToUser: formid.OpenId,
		TplId:  lr.cfg.ReceiveTaskTplId,
		Page:   fmt.Sprintf(RECEIVE_TASK_PAGE, t.EventId, t.ID),
		FormId: formid.FormId,
		Data: &wechat.TaskReceiveTplMsg{
			Keyword1: &template.DataItem{
				Value: t.Name,
				Color: "#008B8B",
			},
			Keyword2: &template.DataItem{
				Value: taskMembers,
				Color: "#008B8B",
			},
			Keyword3: &template.DataItem{
				Value: remindTime,
				Color: "#008B8B",
			},
			Keyword4: &template.DataItem{
				Value: createUser.NickName,
				Color: "#008B8B",
			},
		},
	}
	lr.w.Send(tplMsg)
}

func (lr *ListReminder) Run() {
	c := cron.New()
	c.AddFunc(lr.cfg.ListRemindCron, lr.runRemind)
	c.Start()

	select {
	case <-lr.stop:
		c.Stop()
		return
	}
}

func (lr *ListReminder) runRemind() {
	lr.runTimeRemind(1800, "将于半小时后到期")
	lr.runTimeRemind(86400, "将于一天后到期")
}

func (lr *ListReminder) runTimeRemind(a int64, remark string) {
	start := now.BeginningOfMinute().Unix()
	end := now.EndOfMinute().Unix()
	start += a
	end += a
	holmes.Debug("run remind start: %d end: %d", start, end)

	tasks, err := models.GetRemindTaskList(start, end)
	if err != nil {
		holmes.Error("get remind task list error: %v", err)
		return
	}
	holmes.Debug("get need remind tasks; %v", tasks)
	for i := 0; i < len(tasks); i++ {
		taskMembers, err := models.GetTaskMemberList(tasks[i].ID)
		if err != nil {
			holmes.Error("get task member list error: %v", err)
			continue
		}
		taskTime := time.Unix(tasks[i].RemindTime, 0).Format("2006.01.02 15:04")
		for j := 0; j < len(taskMembers); j++ {
			formid, err := lr.getFormId(taskMembers[j].UserId)
			if err != nil {
				continue
			}
			if formid == nil {
				continue
			}
			tplMsg := &wechat.WechatTplMsg{
				ToUser: formid.OpenId,
				TplId:  lr.cfg.RemindTaskTplId,
				Page:   fmt.Sprintf(REMIND_TASK_PAGE, tasks[i].EventId),
				FormId: formid.FormId,
				Data: &wechat.TaskRemindTplMsg{
					Keyword1: &template.DataItem{
						Value: tasks[i].Name,
						Color: "#008B8B",
					},
					Keyword2: &template.DataItem{
						Value: taskTime,
						Color: "#008B8B",
					},
					Keyword3: &template.DataItem{
						Value: remark,
						Color: "#008B8B",
					},
				},
			}
			lr.w.Send(tplMsg)
		}
	}
}

func (lr *ListReminder) getFormId(userId int64) (*models.TplFormid, error) {
	formIds, err := models.GetTplFormids(userId)
	if err != nil {
		holmes.Error("get tpl form ids error: %v", err)
		return nil, err
	}
	nowTime := time.Now().Unix()
	var expireTime int64 = 0

	defer func() {
		if expireTime != 0 {
			if err = models.DelExpireTplFormids(userId, expireTime); err != nil {
				holmes.Error("del expire tpl form ids erro: %v", err)
			}
		}
	}()
	for i := 0; i < len(formIds); i++ {
		if formIds[i].Expire > nowTime {
			expireTime = formIds[i].Expire
			return &formIds[i], nil
		} else {
			if expireTime < formIds[i].Expire {
				expireTime = formIds[i].Expire
			}
		}
	}

	return nil, nil
}
