package wechat

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
)

type WechatTplMsg struct {
	ToUser          string      `json:"touser"`
	TplId           string      `json:"template_id"`
	Page            string      `json:"page"`
	FormId          string      `json:"form_id"`
	Data            interface{} `json:"data"`
	EmphasisKeyword string      `json:"emphasis_keyword"`
}

type TaskDoneTplMsg struct {
	Keyword1 *template.DataItem `json:"keyword1"`
	Keyword2 *template.DataItem `json:"keyword2"`
	Keyword3 *template.DataItem `json:"keyword3"`
}

type TaskRemindTplMsg struct {
	Keyword1 *template.DataItem `json:"keyword1"`
	Keyword2 *template.DataItem `json:"keyword2"`
	Keyword3 *template.DataItem `json:"keyword2"`
}

type TaskReceiveTplMsg struct {
	Keyword1 *template.DataItem `json:"keyword1"`
	Keyword2 *template.DataItem `json:"keyword2"`
	Keyword3 *template.DataItem `json:"keyword3"`
	Keyword4 *template.DataItem `json:"keyword4"`
}
