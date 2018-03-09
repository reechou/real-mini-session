package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nanjishidu/wechat/small"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/models"
)

func (s *Server) home(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (s *Server) login(c *gin.Context) {
	rsp := &LoginResponse{}
	rsp.Data.Magic = 1

	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	code := c.GetHeader("X-WX-Code")
	encryptedData := c.GetHeader("X-WX-Encrypted-Data")
	iv := c.GetHeader("X-WX-IV")
	appid := c.GetHeader("appid")

	if code == "" || encryptedData == "" || iv == "" || appid == "" {
		holmes.Error("code, encryptedData, iv, appid cannot be nil")
		rsp.Data.Message = ERR_MSG_PARAMS
		return
	}

	wxMini, ok := s.wxMiniMap[appid]
	if !ok {
		appinfo := &models.AppInfo{
			AppId: appid,
		}
		has, err := models.GetAppInfo(appinfo)
		if err != nil {
			holmes.Error("get appinfo error: %v", err)
			rsp.Data.Message = ERR_MSG_SYSTEM
			return
		}
		if !has {
			holmes.Error("cannot found this appid: %s", appinfo)
			rsp.Data.Message = ERR_MSG_SYSTEM
			return
		}
		wxMini = s.addWxMini(appinfo)
	}

	session, err := wxMini.WxMini.GetWxSessionKey(code)
	if err != nil {
		holmes.Error("get wx session key error: %v", err)
		rsp.Data.Message = ERR_MSG_GET_SESSION
		return
	}

	userinfo, err := small.GetWxUserInfo(session.SessionKey, encryptedData, iv)
	if err != nil {
		holmes.Error("get wx user info error: %v", err)
		rsp.Data.Message = ERR_MSG_GET_USER_INFO
		return
	}

	sessionInfo := &models.SessionInfo{
		AppId:  wxMini.AppInfo.ID,
		OpenId: userinfo.OpenId,
	}
	has, err := models.GetSessionInfo(sessionInfo)
	if err != nil {
		holmes.Error("get session info error: %v", err)
		rsp.Data.Message = ERR_MSG_SYSTEM
		return
	}
	if !has {
		sessionInfo.UnionId = userinfo.UnionId
		sessionInfo.NickName = userinfo.NickName
		sessionInfo.AvatarUrl = userinfo.AvatarUrl
		sessionInfo.Gender = int64(userinfo.Gender)
		sessionInfo.Country = userinfo.Country
		sessionInfo.Province = userinfo.Province
		sessionInfo.City = userinfo.City
		sessionInfo.SessionKey = session.SessionKey
		err = models.CreateSessionInfo(sessionInfo)
		if err != nil {
			holmes.Error("create session info error: %v", err)
			rsp.Data.Message = ERR_MSG_SYSTEM
			return
		}
	}

	rsp.Data.Session = &Session{UserInfo: sessionInfo, UserId: sessionInfo.ID}
}

func (s *Server) getUserInfo(c *gin.Context) {
	rsp := &Response{}

	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	appid := c.GetHeader("appid")
	useridStr := c.GetHeader("userid")
	userid, err := strconv.ParseInt(useridStr, 0, 10)
	if err != nil {
		holmes.Error("get user id error: %s", useridStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	holmes.Debug("get user info req: %s %d", appid, userid)

	sessionInfo := &models.SessionInfo{ID: userid}
	has, err := models.GetSessionInfoFromId(sessionInfo)
	if err != nil {
		holmes.Error("get session info from id: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if !has {
		holmes.Error("cannot found this user: %d", userid)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	rsp.Data = &UserInfo{UserInfo: sessionInfo}
}
