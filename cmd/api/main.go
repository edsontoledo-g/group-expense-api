package main

import (
	"github.com/edsontoledo-g/group-expense-api/internal/expenses"
	"github.com/edsontoledo-g/group-expense-api/internal/groups"
	"github.com/edsontoledo-g/group-expense-api/internal/invites"
	"github.com/edsontoledo-g/group-expense-api/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	users.RegisterRoutes(v1)
	expenses.RegisterRoutes(v1)
	groups.RegisterRoutes(v1)
	invites.RegisterRoutes(v1)
	r.Run()
}
