package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/models"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/response"
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

func (s *Server) getListEventDetail(c *gin.Context) {
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

	event := new(models.Event)

	has, err := models.GetEvent(event)
	if err != nil {
		holmes.Error("get event detail error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if !has {
		rsp.Code = ERR_CODE_NOT_FOUND
		rsp.Msg = ERR_MSG_NOT_FOUND
		return
	}
	rsp.Data = event
}

// list event task
func (s *Server) getEventTasks(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	eventIdStr := c.Param("eventid")
	eventId, err := strconv.ParseInt(eventIdStr, 0, 10)
	if err != nil {
		holmes.Error("id str[%s] error", eventIdStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	taskList, err := models.GetTaskInfoDetailList(eventId)
	if err != nil {
		holmes.Error("get task detail list error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = taskList
}

func (s *Server) getTaskDetail(c *gin.Context) {
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

	response := new(TaskDetailRsp)
	response.Task.ID = id

	has, err := models.GetTask(&response.Task)
	if err != nil {
		holmes.Error("get task detail error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if !has {
		rsp.Code = ERR_CODE_NOT_FOUND
		rsp.Msg = ERR_MSG_NOT_FOUND
		return
	}

	response.Members, err = models.GetTaskMemberDetailList(id)
	if err != nil {
		holmes.Error("get task members error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = response
}

// event task members
func (s *Server) getEventTaskMembers(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	eventIdStr := c.Param("eventid")
	eventId, err := strconv.ParseInt(eventIdStr, 0, 10)
	if err != nil {
		holmes.Error("event id str[%s] error", eventIdStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	taskIdStr := c.Param("taskid")
	taskId, err := strconv.ParseInt(taskIdStr, 0, 10)
	if err != nil {
		holmes.Error("task id str[%s] error", taskIdStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	membersRsp := new(EventTaskMembersRsp)
	membersRsp.EventMembers, err = models.GetEventMemberDetailList(eventId)
	if err != nil {
		holmes.Error("get event members error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if taskId != 0 {
		membersRsp.TaskMembers, err = models.GetTaskMemberDetailList(taskId)
		if err != nil {
			holmes.Error("get task members error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
	}
	rsp.Data = membersRsp
}
