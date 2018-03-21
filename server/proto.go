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
