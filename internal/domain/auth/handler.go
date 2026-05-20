package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type verifyPageData struct {
	Title       string
	Message     string
	StatusColor string
	Icon        string
}

type AuthHandler struct {
	s AuthService
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := &AuthInput{
		Email:     req.Email,
		Password:  req.Password,
		Provider:  req.Provider,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	err := h.s.SignUp(input)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an unexpected error occurred"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your registration was successful. We've sent you a verification email. Click the link to activate your account.",
	})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := &AuthInput{
		Email:    req.Email,
		Password: req.Password,
		Provider: req.Provider,
	}
	result, err := h.s.SignIn(input)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.Is(err, ErrAccountNotVerified):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "an unexpected error occurred"})
		}
		return
	}
	res := AuthResponse{
		JWT:       result.JWT,
		ExpiresIn: result.ExpiresIn,
		UserID:    result.UserID,
	}
	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) SignInWithApple(c *gin.Context) {

}

func (h *AuthHandler) Verify(c *gin.Context) {
	token := c.Query("token")
	err := h.s.VerifyUserEmail(token)
	if err != nil {
		h.returnHTMLResult(c, false, "The verification link is invalid or has expired")
		return
	}
	h.returnHTMLResult(c, true, "Your account has been successfully verified!")
}

func (h *AuthHandler) returnHTMLResult(c *gin.Context, success bool, message string) {
	data := verifyPageData{
		Title:       "Verification Error",
		StatusColor: "text-red-500",
		Icon:        "❌",
		Message:     message,
	}
	statusCode := http.StatusBadRequest
	if success {
		data.Title = "All set!"
		data.StatusColor = "text-green-500"
		data.Icon = "✅"
		statusCode = http.StatusOK
	}
	c.HTML(statusCode, "auth_verify.html", data)
}

func NewAuthHandler(s AuthService) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}
