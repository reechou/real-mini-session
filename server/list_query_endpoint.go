package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/models"
)

func (s *Server) getList(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	userIdStr := c.Param("userid")
	userId, err := strconv.ParseInt(userIdStr, 0, 10)
	if err != nil {
		holmes.Error("user id str[%s] error", userIdStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	list, err := models.GetListEvent(userId)
	if err != nil {
		holmes.Error("get list event error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = list
}
