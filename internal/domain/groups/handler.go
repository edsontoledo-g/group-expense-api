package groups

import "github.com/gin-gonic/gin"

type GroupsHandler struct {
	s GroupsService
}

func (_ *GroupsHandler) List(c *gin.Context) {

}

func (_ *GroupsHandler) Get(c *gin.Context) {

}

func (_ *GroupsHandler) Create(c *gin.Context) {

}

func (_ *GroupsHandler) Invite(c *gin.Context) {

}

func NewGroupsHandler(s GroupsService) *GroupsHandler {
	return &GroupsHandler{
		s: s,
	}
}
