package auth

type AuthRequest struct {
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Provider       string  `json:"provider"`
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	ProviderUserID *string `json:"providerUserId"`
}

type AuthResponse struct {
	JWT       string `json:"jwt"`
	ExpiresIn uint   `json:"expiresIn"`
	UserID    uint   `json:"user_id"`
}
