package wechat

import (
	"github.com/reechou/holmes"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

type WechatController struct {
	accessTokenServer core.AccessTokenServer
	wxClient          *core.Client
}

func NewWechatController(appid, secret string) *WechatController {
	wc := &WechatController{}
	wc.accessTokenServer = core.NewDefaultAccessTokenServer(
		appid,
		secret,
		nil)
	wc.wxClient = core.NewClient(wc.accessTokenServer, nil)

	return wc
}

func (wc *WechatController) SendTplMsg(msg *WechatTplMsg) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="

	var result core.Error
	if err = wc.wxClient.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		holmes.Error("send tpl msg error: %v", result)
		err = &result
		return
	}
	return
}
