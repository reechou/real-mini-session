package server

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nanjishidu/wechat/small"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/config"
	"github.com/reechou/real-mini-session/models"
)

type WxMiniInfo struct {
	WxMini  *small.Wx
	AppInfo *models.AppInfo
}

type Server struct {
	cfg *config.Config

	sync.Mutex
	wxMiniMap map[string]*WxMiniInfo
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		cfg:       cfg,
		wxMiniMap: make(map[string]*WxMiniInfo),
	}

	return s
}

func (s *Server) init() {
	models.InitDB(s.cfg)
}

func (s *Server) addWxMini(appinfo *models.AppInfo) *WxMiniInfo {
	s.Lock()
	defer s.Unlock()

	s.wxMiniMap[appinfo.AppId] = &WxMiniInfo{
		WxMini:  small.NewWx(appinfo.AppId, appinfo.Secret),
		AppInfo: appinfo,
	}
	return s.wxMiniMap[appinfo.AppId]
}

func (s *Server) Run() {
	defer holmes.Start(holmes.LogFilePath("./log"),
		holmes.EveryDay,
		holmes.AlsoStdout,
		holmes.DebugLevel).Stop()

	router := gin.Default()

	router.GET("/", s.home)
	router.GET("/mini/login", s.login)
	router.GET("/mini/user", s.getUserInfo)

	holmes.Infoln(router.Run(s.cfg.Host))
}
