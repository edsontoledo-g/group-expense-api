package main

import (
	"github.com/edsontoledo-g/group-expense-api/internal/domain/auth"
	"github.com/edsontoledo-g/group-expense-api/internal/domain/expenses"
	"github.com/edsontoledo-g/group-expense-api/internal/domain/groups"
	"github.com/edsontoledo-g/group-expense-api/internal/domain/invites"
	"github.com/edsontoledo-g/group-expense-api/internal/domain/users"
	"github.com/edsontoledo-g/group-expense-api/internal/infra/db"
	"github.com/edsontoledo-g/group-expense-api/internal/shared/templates"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	tmpl := templates.Load(
		auth.Templates(),
	)
	r.SetHTMLTemplate(tmpl)
	v1 := r.Group("/api/v1")
	pool := db.NewPostgresPool()
	defer pool.Close()
	authModule := auth.NewModule(pool)
	expensesModule := expenses.NewModule(pool)
	groupsModule := groups.NewModule(pool)
	invitesModule := invites.NewModule()
	usersModule := users.NewModule()
	authModule.RegisterRoutes(v1)
	expensesModule.RegisterRoutes(v1)
	groupsModule.RegisterRoutes(v1)
	invitesModule.RegisterRoutes(v1)
	usersModule.RegisterRoutes(v1)
	r.Run()
}
