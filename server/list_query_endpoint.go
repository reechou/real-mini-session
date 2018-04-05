package server

import (
	"net/http"
	"sort"
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

	eventsTasks, err := models.GetEventTaskDetailList(listId)
	if err != nil {
		holmes.Error("get list event task detail error: %v", err)
		holmes.Error("get list event error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	eventMap := make(map[int64]*models.Event)
	for i := 0; i < len(eventsTasks); i++ {
		if _, ok := eventMap[eventsTasks[i].Event.ID]; !ok {
			eventMap[eventsTasks[i].Event.ID] = &eventsTasks[i].Event
			eventMap[eventsTasks[i].Event.ID].TaskNum = 1
		} else {
			eventMap[eventsTasks[i].Event.ID].TaskNum++
		}
	}
	events := make([]*models.Event, 0)
	var keys []int
	for k := range eventMap {
		keys = append(keys, int(k))
		//events = append(events, v)
	}
	sort.Ints(keys)
	for _, k := range keys {
		events = append(events, eventMap[int64(k)])
	}
	rsp.Data = events

	//events, err := models.GetListEvents(listId)
	//if err != nil {
	//	holmes.Error("get list event error: %v", err)
	//	rsp.Code = ERR_CODE_SYSTEM
	//	rsp.Msg = ERR_MSG_SYSTEM
	//	return
	//}
	//rsp.Data = events
}

func (s *Server) getShareEvents(c *gin.Context) {
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

	events, err := models.GetShareEventDetailList(userId)
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

	event := &models.Event{ID: id}

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

// task tag
func (s *Server) getEventTaskTags(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("eventid")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if tags, err := models.GetEventTaskTagList(id); err != nil {
		holmes.Error("get event task tags error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	} else {
		rsp.Data = tags
	}
}

func (s *Server) getTaskTags(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("taskid")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if tags, err := models.GetTaskTagDetailList(id); err != nil {
		holmes.Error("get task tags error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	} else {
		rsp.Data = tags
	}
}

func (s *Server) getEventAndTaskTags(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req EventTaskTagsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	var err error
	result := new(EventTaskTagsRsp)
	if result.EventTags, err = models.GetEventTaskTagList(req.EventId); err != nil {
		holmes.Error("get event task tags error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if result.TaskTags, err = models.GetTaskTagDetailList(req.TaskId); err != nil {
		holmes.Error("get task tags error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	holmes.Debug("result: %v %v", req, result)
	rsp.Data = result
}
