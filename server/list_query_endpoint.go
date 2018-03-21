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

func (s *Server) getListDetail(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	response := new(ListDetailRsp)
	response.List.ID = id

	has, err := models.GetList(&response.List)
	if err != nil {
		holmes.Error("get list detail error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if !has {
		rsp.Code = ERR_CODE_NOT_FOUND
		rsp.Msg = ERR_MSG_NOT_FOUND
		return
	}

	response.Tags, err = models.GetListTags(id)
	if err != nil {
		holmes.Error("get list tags error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = response
}

// list event
func (s *Server) getListEvents(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	listIdStr := c.Param("listid")
	listId, err := strconv.ParseInt(listIdStr, 0, 10)
	if err != nil {
		holmes.Error("list id str[%s] error", listIdStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	events, err := models.GetListEvents(listId)
	if err != nil {
		holmes.Error("get list event error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = events
}
