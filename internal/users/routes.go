package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	h := &UserHandler{}
	user := rg.Group("/me")
	{
		user.GET("", h.GetInformation)
		user.PATCH("/groups/:id/archive", h.ArchiveGroup)
	}
}
