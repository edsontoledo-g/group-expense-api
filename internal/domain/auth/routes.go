package auth

import "github.com/gin-gonic/gin"

func registerRoutes(rg *gin.RouterGroup, h *AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
		auth.POST("/signin/apple", h.SignInWithApple)
		auth.GET("/verify", h.Verify)
	}
}
