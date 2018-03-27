package wechat

import (
	"sync"

	"github.com/reechou/holmes"
)

const (
	DEFAULT_TPL_WORKER_NUM = 1024
)

type WechatTplWorker struct {
	wg sync.WaitGroup
	wc *WechatController

	msgChan   chan *WechatTplMsg
	WorkerNum int

	stop chan struct{}
}

func NewWechatTplWorker(wc *WechatController) *WechatTplWorker {
	wtw := &WechatTplWorker{
		wc:      wc,
		msgChan: make(chan *WechatTplMsg, 10240),
		stop:    make(chan struct{}),
	}
	wtw.WorkerNum = DEFAULT_TPL_WORKER_NUM

	for i := 0; i < wtw.WorkerNum; i++ {
		wtw.wg.Add(1)
		go wtw.runWorker()
	}

	holmes.Debug("wechat tpl worker start..")

	return wtw
}

func (wtw *WechatTplWorker) Stop() {
	close(wtw.stop)
	wtw.wg.Wait()
}

func (wtw *WechatTplWorker) Send(msg *WechatTplMsg) {
	select {
	case wtw.msgChan <- msg:
	case <-wtw.stop:
		return
	}
}

func (wtw *WechatTplWorker) runWorker() {
	for {
		select {
		case msg := <-wtw.msgChan:
			wtw.sendTplMsg(msg)
		case <-wtw.stop:
			wtw.wg.Done()
			return
		}
	}
}

func (wtw *WechatTplWorker) sendTplMsg(msg *WechatTplMsg) {
	err := wtw.wc.SendTplMsg(msg)
	if err != nil {
		holmes.Error("send tpl msg error: %v", err)
	}
}
