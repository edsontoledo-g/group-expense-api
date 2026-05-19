package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// TODO: Move HTML content to a template and use redirect
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
	title := "Verification Error"
	statusColor := "text-red-500"
	icon := "❌"
	statusCode := http.StatusBadRequest
	if success {
		title = "All set!"
		statusColor = "text-green-500"
		icon = "✅"
		statusCode = http.StatusOK
	}
	htmlTemplate := `
	<!DOCTYPE html>
	<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>` + title + `</title>
		<script src="https://cdn.tailwindcss.com"></script>
	</head>
	<body class="bg-gray-50 flex items-center justify-center h-screen">
		<div class="bg-white p-8 rounded-xl shadow-md max-w-md w-full text-center">
			<div class="text-6xl mb-4">` + icon + `</div>
			<h1 class="text-2xl font-bold mb-2 text-gray-800">` + title + `</h1>
			<p class="text-gray-600 mb-6 ` + statusColor + ` font-medium">` + message + `</p>
			<p class="text-xs text-gray-400">You can now close this tab and return to the app.</p>
		</div>
	</body>
	</html>`
	c.Data(statusCode, "text/html; charset=utf-8", []byte(htmlTemplate))
}

func NewAuthHandler(s AuthService) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}
