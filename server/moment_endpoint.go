package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/models"
)

const (
	MOMENT_OPR_TYPE_CREATE_TASK = iota
	MOMENT_OPR_TYPE_DELETE_TASK
	MOMENT_OPR_TYPE_DONE_TASK
	MOMENT_OPR_TYPE_REOPEN_TASK
	MOMENT_OPR_TYPE_ASSIGN_TASK
)

var MOMENT_OPR_MAP = map[int64]string{
	MOMENT_OPR_TYPE_CREATE_TASK: "创建任务",
	MOMENT_OPR_TYPE_DELETE_TASK: "删除任务",
	MOMENT_OPR_TYPE_DONE_TASK:   "完成任务",
	MOMENT_OPR_TYPE_REOPEN_TASK: "重新打开任务",
	MOMENT_OPR_TYPE_ASSIGN_TASK: "指派任务",
}

func (s *Server) createAssignMoment(task *models.Task, oprUser, assignUser int64) {
	user := &models.SessionInfo{ID: assignUser}
	has, err := models.GetSessionInfoFromId(user)
	if err != nil {
		holmes.Error("get session info from id error: %v", err)
		return
	}
	if !has {
		return
	}
	s.createMoment(&models.TaskMoment{
		EventId: task.EventId,
		TaskId:  task.ID,
		UserId:  oprUser,
		OprType: MOMENT_OPR_TYPE_ASSIGN_TASK,
		Detail:  fmt.Sprintf("将任务【%s】指派给：%s", task.Name, user.NickName),
	})
}

func (s *Server) createMoment(tm *models.TaskMoment) {
	tm.OprInfo = MOMENT_OPR_MAP[tm.OprType]
	if err := models.CreateTaskMoment(tm); err != nil {
		holmes.Error("create task moment error: %v", err)
	}
}

func (s *Server) createComment(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req models.TaskComment
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	var err error
	if err = models.CreateTaskComment(&req); err != nil {
		holmes.Error("create task comment error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	user := models.SessionInfo{ID: req.UserId}
	has, err := models.GetSessionInfoFromId(&user)
	if err != nil {
		holmes.Error("get session info from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if !has {
		return
	}
	rsp.Data = models.TaskCommentDetail{TaskComment: req, SessionInfo: user}
}

func (s *Server) getCommentList(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req EventQueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	list, err := models.GetTaskCommentDetailList(req.EventId, req.Offset, req.Num)
	if err != nil {
		holmes.Error("get task comment list error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = list
}

func (s *Server) getMomentList(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req EventQueryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	list, err := models.GetTaskMomentDetailList(req.EventId, req.Offset, req.Num)
	if err != nil {
		holmes.Error("get task moment list error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	rsp.Data = list
}
