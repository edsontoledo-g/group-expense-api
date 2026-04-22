package users

type UserResponse struct {
	ID        string  `json:"id"`
	UserName  string  `json:"username"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	ImageURL  *string `json:"imageUrl,omitempty"`
}
