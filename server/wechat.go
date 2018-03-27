package server

import (
	"github.com/reechou/real-mini-session/wechat"
)

type Wechat struct {
	wc  *wechat.WechatController
	wtw *wechat.WechatTplWorker
}

func NewWechat(appid, secret string) *Wechat {
	w := &Wechat{}

	w.wc = wechat.NewWechatController(appid, secret)
	w.wtw = wechat.NewWechatTplWorker(w.wc)

	return w
}

func (w *Wechat) Send(msg *wechat.WechatTplMsg) {
	w.wtw.Send(msg)
}
