package server

import (
	"github.com/reechou/real-mini-session/models"
)

// session login
type Session struct {
	UserInfo *models.SessionInfo `json:"userInfo"`
	UserId   int64               `json:"user_id"`
}

type LoginResponse struct {
	Magic   int      `json:"F2C224D4-2BCE-4C64-AF9F-A6D872000D1A"`
	Session *Session `json:"session,omitempty"`
	Message string   `json:"message,omitempty"`
}

type UserInfo struct {
	UserInfo *models.SessionInfo `json:"userInfo"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

// list proto
type OprListReq struct {
	List models.List `json:"list"`
	Tags []string    `json:"tags"`
}

type OprTaskReq struct {
	Task    models.Task `json:"task"`
	Members []int64     `json:"members"`
}

type ListDetailRsp struct {
	List models.List      `json:"list"`
	Tags []models.ListTag `json:"tags"`
}

type TaskDetailRsp struct {
	Task    models.Task               `json:"task"`
	Members []models.TaskMemberDetail `json:"members"`
}

type EventTaskMembersRsp struct {
	EventMembers []models.EventMemberDetail `json:"eventMembers"`
	TaskMembers  []models.TaskMemberDetail  `json:"taskMembers"`
}

type DelShareEventReq struct {
	ID      int64 `json:"id"`
	UserId  int64 `json:"userId"`
	EventId int64 `json:"eventId"`
}

type DelEventMemberReq struct {
	ID      int64 `json:"id"`
	UserId  int64 `json:"userId"`
	EventId int64 `json:"eventId"`
}

// form id
type FormIdInfo struct {
	FormId string `json:"formId"`
	Expire int64  `json:"expire"`
}
type SaveFormIdsReq struct {
	UserId  int64        `json:"userId"`
	OpenId  string       `json:"openId"`
	FormIds []FormIdInfo `json:"formIds"`
}
