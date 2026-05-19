package expenses

import (
	"github.com/gin-gonic/gin"
)

type ExpensesHandler struct {
	s ExpensesService
}

func (_ *ExpensesHandler) List(c *gin.Context) {

}

func (_ *ExpensesHandler) Create(c *gin.Context) {

}

func (_ *ExpensesHandler) Delete(c *gin.Context) {

}

func NewExpensesHandler(s ExpensesService) *ExpensesHandler {
	return &ExpensesHandler{
		s: s,
	}
}
