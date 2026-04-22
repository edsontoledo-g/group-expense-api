package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	h := &AuthHandler{}
	auth := rg.Group("/auth")
	{
		auth.GET("/signin", h.SignIn)
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin/apple", h.SignInWithApple)
	}
}
