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
	Wx      *Wechat
}

type Server struct {
	cfg *config.Config

	lr *ListReminder

	sync.Mutex
	wxMiniMap map[string]*WxMiniInfo
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		cfg:       cfg,
		wxMiniMap: make(map[string]*WxMiniInfo),
	}
	s.init()

	return s
}

func (s *Server) init() {
	models.InitDB(s.cfg)

	// list remind
	listAppid := "wx7de642d41cc07693"
	wxMini, ok := s.wxMiniMap[listAppid]
	if !ok {
		appinfo := &models.AppInfo{
			AppId: listAppid,
		}
		has, err := models.GetAppInfo(appinfo)
		if err != nil {
			holmes.Error("get appinfo error: %v", err)
			return
		}
		if !has {
			holmes.Error("cannot found this appid: %s", appinfo)
			return
		}
		wxMini = s.addWxMini(appinfo)
	}
	s.lr = NewListReminder(s.cfg, wxMini.Wx)
}

func (s *Server) addWxMini(appinfo *models.AppInfo) *WxMiniInfo {
	s.Lock()
	defer s.Unlock()

	s.wxMiniMap[appinfo.AppId] = &WxMiniInfo{
		WxMini:  small.NewWx(appinfo.AppId, appinfo.Secret),
		AppInfo: appinfo,
		Wx:      NewWechat(appinfo.AppId, appinfo.Secret),
	}
	return s.wxMiniMap[appinfo.AppId]
}

func (s *Server) Run() {
	defer holmes.Start(holmes.LogFilePath("./log"),
		holmes.EveryDay,
		holmes.AlsoStdout,
		holmes.DebugLevel).Stop()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	//router := gin.Default()

	router.GET("/", s.home)
	router.GET("/mini/login", s.login)
	router.GET("/mini/user", s.getUserInfo)

	router.GET("/test/tpl", s.testTplMsg)

	// list
	router.POST("/list/save", s.saveList)
	router.GET("/list/del/:id", s.delList)
	router.GET("/list/get/:userid", s.getList)
	router.GET("/list/detail/:id", s.getListDetail)
	// event
	router.POST("/event/save", s.saveListEvent)
	router.GET("/event/del/:id", s.delListEvent)
	router.GET("/event/get/:listid", s.getListEvents)
	router.GET("/event/share/get/:userid", s.getShareEvents)
	router.POST("/event/share/del", s.delShareEvent)
	router.GET("/event/detail/:id", s.getListEventDetail)
	router.POST("/event/addmember", s.createListEventMember)
	router.POST("/event/delmember", s.delListEventMember)
	// task
	router.POST("/task/save", s.saveTask)
	router.GET("/task/get/:eventid", s.getEventTasks)
	router.GET("/task/del/:id", s.delTask)
	router.GET("/task/detail/:id", s.getTaskDetail)
	router.GET("/task/done/:id", s.doneTask)
	router.GET("/task/reopen/:id", s.reopenTask)
	// event task members
	router.GET("/event_task/members/:eventid/:taskid", s.getEventTaskMembers)
	// event task tag
	router.POST("/tag/event/save", s.createEventTaskTag)
	router.POST("/tag/task/save", s.createTaskTag)
	router.GET("/tag/event/del/:id", s.delEventTaskTag)
	router.GET("/tag/task/del/:id", s.delTaskTag)
	router.GET("/tag/event/get/:eventid", s.getEventTaskTags)
	router.GET("/tag/event/get/:taskid", s.getTaskTags)
	// form id
	router.POST("/formid/save", s.saveFormIds)

	holmes.Infoln(router.Run(s.cfg.Host))
}
