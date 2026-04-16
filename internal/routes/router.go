package routes

import (
	"github.com/edsontoledo-g/group-expense-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userHandler := handlers.NewUserHandler()
	expensesHandler := handlers.NewExpensesHandler()
	groupsHandler := handlers.NewGroupsHandler()
	invitesHandler := handlers.NewInvitesHandler()
	registerUserRoutes(v1, userHandler)
	registerExpensesRoutes(v1, expensesHandler)
	registerGroupsRoutes(v1, groupsHandler)
	registerInvitesRoutes(v1, invitesHandler)
}
