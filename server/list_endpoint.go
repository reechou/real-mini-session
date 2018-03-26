package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/reechou/holmes"
	"github.com/reechou/real-mini-session/models"
)

func (s *Server) saveList(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req OprListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	var err error
	if req.List.ID == 0 {
		if err = models.CreateList(&req.List); err != nil {
			holmes.Error("create list error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
		listTags := make([]models.ListTag, len(req.Tags))
		for i := 0; i < len(req.Tags); i++ {
			listTags[i].UserId = req.List.UserId
			listTags[i].ListId = req.List.ID
			listTags[i].Name = req.Tags[i]
		}
		err = models.CreateListTags(listTags)
		if err != nil {
			holmes.Error("create list tags error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
	} else {
		if err = models.UpdateList(&req.List); err != nil {
			holmes.Error("update list error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
		tags, err := models.GetListTags(req.List.ID)
		if err != nil {
			holmes.Error("get list tags error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
		newTags := make(map[string]bool)
		for i := 0; i < len(req.Tags); i++ {
			newTags[req.Tags[i]] = true
		}
		oldTags := make(map[string]*models.ListTag)
		for i := 0; i < len(tags); i++ {
			oldTags[tags[i].Name] = &tags[i]
		}
		listTags := make([]models.ListTag, 0)
		for k, _ := range newTags {
			if _, ok := oldTags[k]; !ok {
				listTags = append(listTags, models.ListTag{
					UserId: req.List.UserId,
					ListId: req.List.ID,
					Name:   k,
				})
			}
		}
		if len(listTags) != 0 {
			if err = models.CreateListTags(listTags); err != nil {
				holmes.Error("create list tags error: %v", err)
				rsp.Code = ERR_CODE_SYSTEM
				rsp.Msg = ERR_MSG_SYSTEM
				return
			}
		}
		for k, v := range oldTags {
			if _, ok := newTags[k]; !ok {
				if err = models.DelListTag(v); err != nil {
					holmes.Error("del list tag error: %v", err)
					rsp.Code = ERR_CODE_SYSTEM
					rsp.Msg = ERR_MSG_SYSTEM
					return
				}
			}
		}
	}
}

func (s *Server) delList(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.DelList(&models.List{ID: id}); err != nil {
		holmes.Error("del list from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

// list event
func (s *Server) saveListEvent(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req models.Event
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	var err error
	if req.ID == 0 {
		if err = models.CreateEvent(&req); err != nil {
			holmes.Error("create event error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}

		eventMember := &models.EventMember{
			EventId: req.ID,
			UserId:  req.OwnerUserId,
		}
		if err = models.CreateEventMember(eventMember); err != nil {
			holmes.Error("create event member error: %v", err)
		}
	} else {
		if err = models.UpdateEvent(&req); err != nil {
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
	}
	rsp.Data = req
}

func (s *Server) delListEvent(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.DelEvent(&models.Event{ID: id}); err != nil {
		holmes.Error("del event from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

func (s *Server) delShareEvent(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req DelShareEventReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	holmes.Debug("del share event req: %+v", req)

	var err error
	if err = models.DelShareEvent(&models.ShareEvent{ID: req.ID}); err != nil {
		holmes.Error("del share event from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	if err = models.DelEventMemberFromUserEvent(&models.EventMember{EventId: req.EventId, UserId: req.UserId}); err != nil {
		holmes.Error("del event member from user-event error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

func (s *Server) createListEventMember(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req models.EventMember
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}

	err := models.CreateEventMember(&req)
	if err != nil {
		holmes.Error("create event member error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
	shareEvent := &models.ShareEvent{
		UserId:  req.UserId,
		EventId: req.EventId,
	}
	if err = models.CreateShareEvent(shareEvent); err != nil {
		holmes.Error("create share event error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

func (s *Server) delListEventMember(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.DelEventMember(&models.EventMember{ID: id}); err != nil {
		holmes.Error("del event member from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

// list event task
const (
	TASK_STATUS_NOT_DONE = iota
	TASK_STATUS_DONE
)

func (s *Server) saveTask(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	var req OprTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		holmes.Error("bind json error: %v", err)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if req.Task.Date != "" || req.Task.Time != "" {
		remindT, err := now.Parse(fmt.Sprintf("%s %s", req.Task.Date, req.Task.Time))
		if err != nil {
			holmes.Error("parse remind data and time error: %v", err)
		} else {
			req.Task.RemindTime = remindT.Unix()
		}
	}
	var err error
	if req.Task.ID == 0 {
		err = models.CreateTask(&req.Task)
		if err != nil {
			holmes.Error("create task error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			return
		}
		taskMembers := make([]models.TaskMember, len(req.Members))
		for i := 0; i < len(req.Members); i++ {
			taskMembers[i].TaskId = req.Task.ID
			taskMembers[i].UserId = req.Members[i]
		}
		err = models.CreateTaskMembers(taskMembers)
		if err != nil {
			holmes.Error("create task members error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
	} else {
		err = models.UpdateTask(&req.Task)
		if err != nil {
			holmes.Error("update task error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
		members, err := models.GetTaskMemberList(req.Task.ID)
		if err != nil {
			holmes.Error("get task member list error: %v", err)
			rsp.Code = ERR_CODE_SYSTEM
			rsp.Msg = ERR_MSG_SYSTEM
			return
		}
		newMembers := make(map[int64]bool)
		for i := 0; i < len(req.Members); i++ {
			newMembers[req.Members[i]] = true
		}
		oldMembers := make(map[int64]*models.TaskMember)
		for i := 0; i < len(members); i++ {
			oldMembers[members[i].UserId] = &members[i]
		}
		newAddMembers := make([]models.TaskMember, 0)
		for k, _ := range newMembers {
			if _, ok := oldMembers[k]; !ok {
				newAddMembers = append(newAddMembers, models.TaskMember{
					TaskId: req.Task.ID,
					UserId: k,
				})
			}
		}
		if len(newAddMembers) != 0 {
			err = models.CreateTaskMembers(newAddMembers)
			if err != nil {
				holmes.Error("create task members error: %v", err)
				rsp.Code = ERR_CODE_SYSTEM
				rsp.Msg = ERR_MSG_SYSTEM
				return
			}
		}
		for k, v := range oldMembers {
			if _, ok := newMembers[k]; !ok {
				if err := models.DelTaskMember(v); err != nil {
					holmes.Error("del task members error: %v", err)
					rsp.Code = ERR_CODE_SYSTEM
					rsp.Msg = ERR_MSG_SYSTEM
					return
				}
			}
		}
	}
}

func (s *Server) delTask(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.DelTask(&models.Task{ID: id}); err != nil {
		holmes.Error("del task from id error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

func (s *Server) doneTask(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.UpdateTaskStatus(&models.Task{ID: id, Status: TASK_STATUS_DONE}); err != nil {
		holmes.Error("update task error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}

func (s *Server) reopenTask(c *gin.Context) {
	rsp := &Response{}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 10)
	if err != nil {
		holmes.Error("del id str[%s] error", idStr)
		rsp.Code = ERR_CODE_PARAMS
		rsp.Msg = ERR_MSG_PARAMS
		return
	}
	if err = models.UpdateTaskStatus(&models.Task{ID: id, Status: TASK_STATUS_NOT_DONE}); err != nil {
		holmes.Error("update task error: %v", err)
		rsp.Code = ERR_CODE_SYSTEM
		rsp.Msg = ERR_MSG_SYSTEM
		return
	}
}
